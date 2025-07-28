
header = '''syntax="proto3";
package pb;
option go_package="./;myproto";
'''

message_template = '''
message Request{index} {{
    string f1 = 1; // field1
    uint32 f2 = 2; // field2
    bool f3 = 3; // field3
    bytes f4 = 4; // field4
    uint64 f5 = 5; // field5
    int32 f6 = 6; // field6
    int64 f7 = 7; // field7
}}'''

with open("/home/zhangjie/github/gogo_protobuf/test/memleak/huge.proto", "w") as f:
    f.write(header)
    for i in range(1, 2223):
        f.write(message_template.format(index=i))

print("Generated huge.proto with 2222 messages.")
