import os


proto_path = "proto"
csfile_path = "client/Assets/Scripts/Net/ProtoDic.cs"
gofile_path = "server/msg/msg.go"



def genGolangfile(protos):
    fileContent = ""
    fileContent += (
    'package msg'
    '\n'
    'import (\n'
    '	"github.com/name5566/leaf/network/protobuf"\n'
    ')'
    '\n\n'
    'var (\n'
    '	Processor = protobuf.NewProcessor()\n'
    ')\n'
    '\n'
    'func init() {'
     '	// 这里我们注册 protobuf 消息)\n'
     '    Processor.SetByteOrder(true)\n'
    )


    for index in range(len(protos)):
        fileContent += '' \
                       '    Processor.Register(&'
        fileContent += protos[index]
        fileContent += '{})\n'

    fileContent += ("\n}")

    fo = open(gofile_path, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()


def loadProto():
    protos = []
    for dir in os.listdir(proto_path):  # 遍历当前目录所有文件和目录
        child = os.path.join(proto_path, dir)  # 加上路径，否则找不到
        if os.path.isdir(child):  # 如果是目录，则继续遍历子目录的文件
            for file in os.listdir(child):
                if os.path.splitext(file)[1] == '.proto':  # 分割文件名和文件扩展名，并且扩展名为'proto'
                    file = os.path.join(child, file)  # 同样要加上路径
                    f = open(file, 'r')
                    f_info = f.readlines()  # 获取文件所有内容
                    f.close()
                    fileProtos = findProtos(f_info)
                    protos = protos +fileProtos
        elif os.path.isfile(child):  # 如果是文件，则直接判断扩展名
            if os.path.splitext(child)[1] == '.proto':
                f = open(child, 'r')
                f_info = f.readlines()
                f.close()
                fileProtos = findProtos(f_info)
                protos = protos + fileProtos

    return  protos

def findProtos(lines):
    fileProtos = []
    for line in lines:
        if "message" in line:
            line = line.strip()
            fields = line.split(' ')
            fileProtos.append(fields[1])
        # else:
        #     continue
    return  fileProtos



def genCSfile(protos):
    fileContent = ""
    fileContent += ('\n'
                    'using Google.Protobuf;\n'
                    'using Msg;\n'
                    'using System;\n'
                    'using System.Collections.Generic;\n\n'
                    'namespace Proto\n'
                    '{\n'
                    '   public class ProtoDic\n'
                    '   {\n'
                    '       private static List<int> _protoId = new List<int>\n'
                    '       {\n')

    for index in range(len(protos)):
        fileContent = fileContent + '''            ''' + str(index) + ''',\n'''

    fileContent += '        };\n\n      private static List<Type>_protoType = new List<Type>\n      {\n'

    for index in range(len(protos)):
        fileContent = fileContent + '''            typeof(''' + protos[index] + '''),\n'''

    fileContent += ('       };\n\n'
                    '       private static readonly Dictionary<RuntimeTypeHandle, MessageParser> Parsers = new Dictionary<RuntimeTypeHandle, MessageParser>()\n'
                    '       {\n')

    for index in range(len(protos)):
        fileContent = fileContent + '''            {typeof(''' + protos[index] + ''').TypeHandle,''' + protos[index] + '''.Parser },\n'''

    fileContent += ('       };\n\n        public static MessageParser GetMessageParser(RuntimeTypeHandle typeHandle)\n'
                   '        {\n'
                   '            MessageParser messageParser;\n'
                   '            Parsers.TryGetValue(typeHandle, out messageParser);\n'
                   '            return messageParser;\n'
                   '        }\n'
                   '\n'
                   '        public static Type GetProtoTypeByProtoId(int protoId)\n'
                   '        {\n'
                   '            int index = _protoId.IndexOf(protoId);\n'
                   '            return _protoType[index];\n'
                   '        }\n'
                   '\n'
                   '        public static int GetProtoIdByProtoType(Type type)\n'
                   '        {\n'
                   '            int index = _protoType.IndexOf(type);\n'
                   '            return _protoId[index];\n'
                   '        }\n'
                   '\n'
                   '        public static bool ContainProtoId(int protoId)\n'
                   '        {\n'
                   '            if(_protoId.Contains(protoId))\n'
                   '            {\n'
                   '                return true;\n'
                   '            }\n'
                   '            return false;\n'
                   '        }\n'
                   '\n'
                   '        public static bool ContainProtoType(Type type)\n'
                   '        {\n'
                   '            if(_protoType.Contains(type))\n'
                   '            {\n'
                   '                return true;\n'
                   '            }\n'
                   '            return false;\n'
                   '        }\n'
                   '    }\n'
                   '}')

    fo = open(csfile_path, "wb")
    fo.write(fileContent.encode('utf-8'))
    fo.close()


if __name__ == "__main__":
    protos = loadProto()
    genCSfile(protos)
    genGolangfile(protos)