#coding=utf-8
import os
import re

import re
#test_str = 'NumberInt(1),NumberInt(2),NumberInt(3)'
#regex = r"(NumberInt\( *)(\d)( \))"
#result = re.sub(regex, lambda x:x.group(2), test_str)
def main():
    fileDir = "../proto"
    #print(os.getcwd())
    f = open("../msg/msg.go", 'r',encoding='utf-8')
    all_the_lines = f.readlines()
    
    f.seek(0)
    cpp = open("cpp.txt", mode="w",encoding='utf-8')
    for line in all_the_lines:
        if len(line)>0 and 0 < line.find('RegisterMessage') and 0 < line.find('{}'):
            line = line.replace("RegisterMessage(&protoMsg.","registerMsg(typeid(go::") 
            line = line.replace("{})",").name());")
            cpp.write(line)
    f.close()
    cpp.close()
    print('FINISH')    



if __name__ == '__main__':
    main()