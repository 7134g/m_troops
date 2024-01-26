package setting

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"reflect"
	"time"
)

var Global *ServiceSetting

type ServiceSetting struct {
	rds *redis.Client

	rdsKey string // redis setting key
	subKey string // redis sub key

	allSetting // 所有配置
}

// Monitor 用于监听是否更新配置信息, 建议采用 goroutine
func Monitor() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Monitor recover %v\n", err)
			// 意外退出后，重启
			time.Sleep(time.Minute)
			Monitor()
		}
	}()

	sub := Global.subscribe()
	defer func(sub *redis.PubSub) {
		_ = sub.Close()
	}(sub)

	for {
		select {
		case <-sub.Channel():
			//fmt.Println("接收到redis更新通知")
			err := Global.LoadRedisData(context.Background())
			if err != nil {
				log.Println("redis key error", err)
			}
		}
	}
}

func Initialize(rds *redis.Client, env string) error {
	s := ServiceSetting{
		rds:    rds,
		rdsKey: fmt.Sprintf("setting:%s", env),
		subKey: fmt.Sprintf("setting:pub_sub:%s", env),
	}
	if err := s.ParseTag(); err != nil {
		return err
	}
	if err := s.LoadRedisData(context.Background()); err != nil {
		return err
	}

	Global = &s
	return nil
}

// ParseTag 通过反射读取tag中包含的信息
func (s *ServiceSetting) ParseTag() error {
	serviceSettingType := reflect.TypeOf(*s)
	serviceSettingValue := reflect.ValueOf(s).Elem()

	allSettingType, ok := serviceSettingType.FieldByName("allSetting")
	if !ok {
		return errors.New("allSetting is not exist")
	}
	allSettingValue := serviceSettingValue.FieldByName("allSetting")

	for i := 0; i < allSettingType.Type.NumField(); i++ {
		structName := allSettingType.Type.Field(i).Name
		recordType, ok := allSettingType.Type.FieldByName(structName)
		if !ok {
			return errors.New(structName + " is not exist")
		}
		recordValue := allSettingValue.FieldByName(structName)

		// recordType.Type 指向 levelComment 这一级
		srcValue := recordValue
		dstType := recordType.Type
		for i := 0; i < srcValue.NumField(); i++ {
			name := srcValue.Type().Field(i).Name
			dstStructType, ok := dstType.FieldByName(name)
			if !ok {
				continue
			}
			if dstType.Kind() != reflect.Struct {
				continue
			}

			//  标签:结构体成员名
			tags := map[string]string{
				"json":        "Field",
				"default":     "DefaultValue",
				"description": "Description",
				"category":    "Category",
			}
			for tag, attr := range tags {
				tagValue := dstStructType.Tag.Get(tag)
				if tagValue == "" {
					return errors.New(fmt.Sprintf("%s %s is none", name, tag))
				}

				//tagValue := dstStructType.Tag.Get(tag)
				dstStructTypeToFieldType, ok := dstStructType.Type.FieldByName(attr)
				if !ok {
					continue
				}
				newValue, err := convertToFieldType(tagValue, dstStructTypeToFieldType)
				if err != nil {
					return errors.New(fmt.Sprintf("%s %s --> type error: %v", name, tag, err))
				}

				dstStructValue := srcValue.FieldByName(name)                                     // 获取 Level0SingleCommentMax 值
				dstStructValueToFieldValue := dstStructValue.FieldByName(attr)                   // 获取 Level0SingleCommentMax 中的字段值
				if dstStructValueToFieldValue.IsValid() && dstStructValueToFieldValue.CanSet() { // 判断是否有效且可设置值
					dstStructValueToFieldValue.Set(newValue)

				}
			}
		}
	}

	return nil
}

// LoadRedisData 加载 redis 配置, 通过redis标签读取key，从 HGetAll 的 map 找到数据值， 通过反射将对应的成员中的属性 Value 复制为 redis的值
func (s *ServiceSetting) LoadRedisData(ctx context.Context) error {
	var (
		tag       = "redis"
		attr      = "Value"
		valueType = "DefaultValue"
	)

	m, err := s.rds.HGetAll(ctx, s.rdsKey).Result()
	if err != nil {
		return err
	}
	if len(m) == 0 {
		return errors.New("redis key is empty")
	}

	serviceSettingType := reflect.TypeOf(*s)
	serviceSettingValue := reflect.ValueOf(s).Elem()

	allSettingType, ok := serviceSettingType.FieldByName("allSetting")
	if !ok {
		return errors.New("allSetting is not exist")
	}
	allSettingValue := serviceSettingValue.FieldByName("allSetting")

	for i := 0; i < allSettingType.Type.NumField(); i++ {
		structName := allSettingType.Type.Field(i).Name
		recordType, ok := allSettingType.Type.FieldByName(structName)
		if !ok {
			return errors.New(structName + " is not exist")
		}
		recordValue := allSettingValue.FieldByName(structName)

		srcValue := recordValue
		dstType := recordType.Type
		for i := 0; i < srcValue.NumField(); i++ {
			name := srcValue.Type().Field(i).Name
			dstStructType, ok := dstType.FieldByName(name)
			if !ok {
				return errors.New("check you struct name, can not find " + name)
			}
			if dstType.Kind() != reflect.Struct {
				continue
			}

			// 判断类型
			tagValue := dstStructType.Tag.Get(tag)
			if tagValue == "" {
				// 若没有设置 redis tag内容，取json值
				tagValue = dstStructType.Tag.Get("json")
			}
			value, ok := m[tagValue]
			if !ok {
				// redis中无该数据
				value = emptyTypeValue
			}
			dstStructTypeToFieldType, ok := dstStructType.Type.FieldByName(valueType)
			if !ok {
				return errors.New(fmt.Sprintf("%s struct cannot find %s", name, valueType))
			}
			newValue, err := convertToFieldType(value, dstStructTypeToFieldType)
			if err != nil {
				// 找不到该类型
				return errors.New(fmt.Sprintf("%s %s --> type error: %v", name, tag, err))
			}

			// 赋值
			dstStructValue := srcValue.FieldByName(name)
			dstStructValueToFieldValue := dstStructValue.FieldByName(attr)
			if dstStructValueToFieldValue.IsValid() && dstStructValueToFieldValue.CanSet() {
				dstStructValueToFieldValue.Set(newValue)

			}
		}

	}

	return nil
}

func (s *ServiceSetting) GetAllSetting() (map[string]record, error) {
	b, err := json.Marshal(s.allSetting)
	if err != nil {
		return nil, err
	}

	var m map[string]record
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func (s *ServiceSetting) UpdateRecord(ctx context.Context, data fieldInterface) error {
	if err := s.updateRedis(ctx, data); err != nil {
		return err
	}

	if err := s.LoadRedisData(ctx); err != nil {
		return err
	}

	if err := s.sendNotify(ctx); err != nil {
		return err
	}

	return nil
}

func (s *ServiceSetting) updateRedis(ctx context.Context, data fieldInterface) error {
	return s.rds.HSet(ctx, s.rdsKey, data.GetField(), data.GetValue()).Err()
}

func (s *ServiceSetting) deleteRedis(ctx context.Context, data fieldInterface) error {
	return s.rds.HDel(ctx, s.rdsKey, data.GetField()).Err()
}

func (s *ServiceSetting) subscribe() *redis.PubSub {
	sub := s.rds.Subscribe(context.Background(), s.rdsKey)
	return sub
}

func (s *ServiceSetting) sendNotify(ctx context.Context) error {
	return s.rds.Publish(ctx, s.rdsKey, "update").Err()
}
