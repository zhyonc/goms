# goms
TW MapleStory Server Demo.  
Only implemented the login server.  
The channel server need much time to overcome... 
## Runtime
- MapleVersion: 240.3
- Database: mongodb
- Golang version at least 1.21 to support slog
- Client must resovle old domain to server ip  
```echo 127.0.0.1 tw.login.maplestory.gamania.com >> C:\Windows\System32\drivers\etc\hosts```
## IDB
- [KMS1029](https://forum.ragezone.com/threads/leak-maplestory-korea-test-v-1029-raw-exe-w-debuggables.1100141/)
## Packet
|Header|AESOFB|Note|
|:---:|:---:|:---:|
|4 Bytes|Any Bytes|Except for the first packet|
### Decode
PacketLen = (Header[0]+Header[1]*0x100) ^ (Header[2]+Header[3]*0x100)
### Encode
- sVersion = (^clientVer >> 8 & 0xFF) | ((^clientVer << 8) & 0xFF00)
- a = int(sendIV[3])
- a |= int(sendIV[2])<<8
- a ^= sVersion
- b = ((PacketLen << 8) & 0xFF00) | (PacketLen >> 8)
- c = a ^ b
- Header = [a>>8, a, c>>8, b]
### PacketFormat
|Opcode|Data|Note|
|:---:|:---:|:---:|
|2 Bytes|Any Bytes|Except for the first packet|
### The First Packet
|Name|PacketLen|Version|MinorVersionLen|MinorVersion|RecvIV|SendIV|Region|IsTestServer|Note|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|Connect|2 Bytes|2 Bytes|2 Bytes|1 Byte|4 Bytes|4 Bytes|1 Byte|1 Byte|The first packet|