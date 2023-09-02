import socket
import sys

HOST = 'localhost'
PORT = 9869  
BufferSize = 1024
Address = (HOST, PORT)

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(Address) 

while True:
    print("============================================")
    print("Send message to the example accepter ") 
    print("or input \"quit\" to kill the connection: ")
    data = sys.stdin.readline().strip('\n') 
    if data == "quit":
        client.close()
        break
    client.send(str.encode(data))  
    data = client.recv(BufferSize)  
    print()
    print("response arrived: " + bytes.decode(data))
    print("============================================\n")
