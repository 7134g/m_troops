### 例子

- ./duplication.exe -s "E:\source" -d "E:\destination"

  - 第一步，计算 -s 路径下的所有文件md5，并记录为[src]，筛选出自身出现的重复文件[dup]
  - 第二步，再计算 -d 路径下的所有文件md5，筛选出重复文件[dup]
  - 第三步，输出一个树形结构数据


如下：

- 若开启 -r 参数，只保留下第一步所记录的文件，其下 [dup] 将被删除
- 在启用-r时，设置 -k 值为 "\source", 将只删除包含该值的路径

- `./duplication.exe -s "E:\source" -d "E:\destination" -r -k "\source"`

```text
  [dup]   E:\source\a.txt        删除
  [dup]   E:\destination\a.txt   不删除
```
