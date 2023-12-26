#ac自动机
import ahocorasick as ah
import config
import psutil
import os

ac = ah.Automaton()
keywords = config.keywords

# 利用 add_word方法 将关键词加入自动机！
# 该方法必须包含2个参数，第一个参数是检索词，第二个参数可以任意。
for i, keyword in enumerate(keywords):
    ac.add_word(' ' + keyword + ' ', (i, keyword))

ac.make_automaton()

# 开始查找
# 该方法 匹配所有字符串
print(ac.iter(config.desc.lower()))

for row in ac.iter(config.desc.lower()):
    print((row[1]))
# print(list(ac.iter(config.desc.lower())))
# print(list(ac.iter(' ')))