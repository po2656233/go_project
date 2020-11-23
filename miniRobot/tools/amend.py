#coding=utf-8
import os
import re
number = 0
def listFiles(dirPath):
    fileList = []
    for root, dirs, files in os.walk(dirPath):
        for fileObj in files:
            if fileObj.endswith(".proto"):
                fileList.append(os.path.join(root, fileObj))
    return fileList
def dealSub(s1):
    global number 
    number = number+1
    return '= '+str(number)+';'
    

def deal(s):
    m1  =  '='
    m2  =  ';'
    patN = re.compile (m1 + '(.*?)' + m2)
    data = s.group()
    resultN  =  patN.findall(str(data))
    if 0 == len(resultN):
        return data
    str1 =resultN[0].replace(' ',resultN[0])
    if int(str1) == 0:
        return data
    global number 
    number = 0
    word = re.sub(patN, dealSub, str(data))
    return word
    
def main():
    fileDir = "../msg/proto"
    fileList = listFiles(fileDir)
    #关键字1,2(修改引号间的内容)
    w1  =  '{'
    w2  =  '}'

    for fileObj in fileList:
        #调整协议编码
        f  =  open ( fileObj, 'r+' ,encoding='utf-8')
        buff = f.read()
        pat = re.compile (w1 + '(.*?)' + w2,re.S)
        result  =  pat.findall(buff)
        newNUMS = re.sub(pat, deal, buff)
        
        #包名本地化
        f.seek(0)
        f.truncate()
        all_the_lines = newNUMS.split('\n')
        for line in all_the_lines:
            if len(line)>0:#当改行为空，表明已经读取到文件末尾，退出循环
                content = line.split(' ')#因为每行有三个TAB符号分开的数字，将它们分开
                if len(content)>0 and content[0]=='package':
                    f.write("package go;\n")
                else:
                    f.write(line+'\n')
        f.close()
    
    fd  =  open ( "../msg/msg.go", 'r+' ,encoding='utf-8')
    buffd = fd.read()
    patd = re.compile (w1 + '(.*?)' + w2,re.S)
    resultd  =  pat.findall(buffd)
    newNUMSd = re.sub(pat, deal, buffd)
    
    #包名本地化
    fd.seek(0)
    fd.truncate()
    all_the_linesd = newNUMSd.split('\n')
    for line in all_the_linesd:
        line = line.replace("server/msg/go","miniRobot/msg/go")
        fd.write(line+'\n')    
    fd.close()
    
    
if __name__ == '__main__':
    main()