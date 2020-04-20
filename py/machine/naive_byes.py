import numpy as np
from sklearn.externals import joblib
from sklearn.metrics import accuracy_score
from sklearn.model_selection import train_test_split
from sklearn.naive_bayes import MultinomialNB
from sklearn.utils import shuffle


# 分词
def split_word(words):
    # 将单词以空格划分
    words = words.split()
    # 去除单词长度小于2的无用单词
    return [tok.lower() for tok in words if len(tok) > 2]


# 去列表中重复元素，并以列表形式返回
def creat_voval_table(data: list):
    voval_table = set({})
    # 去重复元素，取并集
    for document in data:
        voval_table = voval_table | set(document)
    return list(voval_table)


# 统计每一文档（或邮件）在单词表中出现的次数，并以列表形式返回
def word_frequency(vocab_table: list, all_word: list) -> list:
    # 创建0向量，其长度为词汇表长度
    doc_word_frequency = [0] * len(vocab_table)
    # 统计相应的词汇出现的数量
    for word in all_word:
        if word in vocab_table:
            doc_word_frequency[vocab_table.index(word)] += 1
    return doc_word_frequency


# 处理数据
def data_deal(x, y):
    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size = 0.25)
    x_train, x_test = np.array(x_train), np.array(x_test)
    y_train, y_test = np.array(y_train), np.array(y_test)
    return x_train, x_test, y_train, y_test


def load_modele(file_name, origin_data_lable=None, test_data=None):
    print('重新加载')
    new_bys = joblib.load(file_name)
    # 测试数据，返回结果
    y_pred = new_bys.predict(test_data)

    # 输出
    print("正确值：{0}".format(origin_data_lable))
    print("预测值：{0}".format(y_pred))
    print("准确率：%f%%" % (accuracy_score(origin_data_lable, y_pred) * 100))

    return new_bys


def dump_modele(file_name, modele):
    joblib.dump(modele, file_name)


def main():
    doc_list, category, x = [], [], []
    for i in range(1, 26):
        # 读取第i篇垃圾文件，并以列表形式返回
        all_word = split_word(open('./email/spam/{0}.txt'.format(i)).read())
        doc_list.append(all_word)
        category.append(1)  # 标记文档为垃圾文档

        # 读取第i篇非垃圾文件，并以列表形式返回
        all_word = split_word(open('./email/ham/{0}.txt'.format(i)).read())
        doc_list.append(all_word)
        category.append(0)  # 标记文档为非垃圾文档

    doc_list, category = shuffle(doc_list, category)
    vocab_table = creat_voval_table(doc_list)  # 创建词汇表，不重复
    # 将数据向量化，词频
    for all_word in doc_list:
        x.append(word_frequency(vocab_table, all_word))
    # 分割为训练集和测试集
    x_train, x_test, y_train, y_test = data_deal(x, category)
    # 用训练集去训练数据
    bys_modele = MultinomialNB()
    bys_modele.fit(x_train, y_train)

    # 测试数据，返回测试结果
    y_pred = bys_modele.predict(x_test)

    # 输出
    print("正确值：{0}".format(y_test))
    print("预测值：{0}".format(y_pred))
    print("准确率：%f%%" % (accuracy_score(y_test, y_pred) * 100))

    # file_name = 'm1.pkl'
    # dump_modele(file_name, bys_modele)
    # load_modele(file_name, y_test, x_test)


if __name__ == '__main__':
    main()
