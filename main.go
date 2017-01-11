package main

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go main.go

//sys   FormatMessage(flags uint32, source syscall.Handle, messageID uint32, languageID uint32, buffer *byte, bufferSize uint32, arguments uintptr) (numChars uint32, err error) = kernel32.FormatMessageW
//sys   MidiOutOpen(lphmo *uint32, deviceID uint, callback uint32, callbackIns uint32, flags uint32)(res WinmmResult, err error)= winmm.midiOutOpen
//sys   MidiOutGetNumDevs() (ret uint, err error) = winmm.midiOutGetNumDevs
//sys   MidiOutShortMsg(hmo uint32, dwmsg uint32)(res WinmmResult, err error)=winmm.midiOutShortMsg
//sys  midiOutGetDevCaps(id uint,data *byte, len uint)(res WinmmResult,err error)=winmm.midiOutGetDevCaps
import(
//	"syscall"
//	"unsafe"
	"log"
"strconv"
"os"
"bufio"
"fmt"
)
/*
"midiConnect"
"midiDisconnect"
"midiInUnprepareHeader"
"midiOutCacheDrumPatches"
"midiOutCachePatches"
"midiOutClose"
"midiOutGetDevCapsA"
"midiOutGetDevCapsW"
"midiOutGetErrorTextA"
"midiOutGetErrorTextW"
"midiOutGetID"
"midiOutGetNumDevs"
"midiOutGetVolume"
"midiOutLongMsg"
"midiOutMessage"
"midiOutOpen"
"midiOutPrepareHeader"
"midiOutReset"
"midiOutSetVolume"
"midiOutShortMsg"
"midiOutUnprepareHeader"
*/

type WinmmResult int
const(
NoError WinmmResult	=0
UnspecifiedError	=1
BadDeviceId=	2
NotEnabled=	3
AlreadyAllocated	=4
InvalidHandle=	5
NoDriver=	6
MemoryAllocationError=	7
NotSupported=	8
BadErrorNumber=	9
InvalidFlag=	10
InvalidParameter=	11
HandleBusy=	12
InvalidAlias=	13
BadRegistryDatabase=	14
RegistryKeyNotFound=	15
RegistryReadError=	16
RegistryWriteError=	17
RegistryDeleteError=	18
RegistryValueNotFound=	19
NoDriverCallback=	20
MoreData=	21
WaveBadFormat=	32
WaveStillPlaying=	33
WaveHeaderUnprepared=	34
WaveSync=	35
AcmNotPossible=	512
AcmBusy=	513
AcmHeaderUnprepared=	514
AcmCancelled=	515
MixerInvalidLine=	1024
MixerInvalidControl=	1025
MixerInvalidValue=	1026
)

type midioutcaps struct{
     wMid            uint32   
     wPid            uint32   
   vDriverVersion    uint32
    szPname          []byte   
     wTechnology     uint32   
     wVoices         uint32   
     wNotes          uint32   
     wChannelMask    uint32   
    dwSupport        uint32   
}





func main(){

b,_:=MidiOutGetNumDevs()
dup:=" "
for i:=0;i<int(b);i++{
 dup+=fmt.Sprintf("%d ",i)
}
log.Print("select "+dup+" : ")
s,_,_:=bufio.NewReader(os.Stdin).ReadLine()
sele,_:=strconv.Atoi(string(s))
log.Print(fmt.Sprintf("select is %d",sele))
h:=uint32(0)
	v,_:=MidiOutOpen(&h,uint(sele),0,0,0)
log.Print(v)
ret,_:=MidiOutShortMsg(h,0x00703C90)
log.Print(ret)
log.Print(h)
bufio.NewReader(os.Stdin).ReadString('\n')
}