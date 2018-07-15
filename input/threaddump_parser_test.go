package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseThreaddump(t *testing.T) {
	parser := CreateInput(`Full thread dump Java HotSpot(TM) 64-Bit Server VM (25.20-b23 mixed mode):

"nioEventLoopGroup-2-1" #85 prio=10 os_prio=0 tid=0x00007f648976f000 nid=0x424c runnable [0x00007f63e76f7000]
   java.lang.Thread.State: RUNNABLE
	at sun.nio.ch.EPollArrayWrapper.epollWait(Native Method)
	at sun.nio.ch.EPollArrayWrapper.poll(EPollArrayWrapper.java:269)
	at sun.nio.ch.EPollSelectorImpl.doSelect(EPollSelectorImpl.java:79)
	at sun.nio.ch.SelectorImpl.lockAndDoSelect(SelectorImpl.java:86)
	- locked <0x000000073b81b6e8> (a io.netty.channel.nio.SelectedSelectionKeySet)
	- locked <0x000000073b8197c0> (a java.util.Collections$UnmodifiableSet)
	- locked <0x000000073b8196e8> (a sun.nio.ch.EPollSelectorImpl)
	at sun.nio.ch.SelectorImpl.select(SelectorImpl.java:97)
	at io.netty.channel.nio.NioEventLoop.select(NioEventLoop.java:622)
	at io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:310)
	at io.netty.util.concurrent.SingleThreadEventExecutor$2.run(SingleThreadEventExecutor.java:116)
	at io.netty.util.concurrent.DefaultThreadFactory$DefaultRunnableDecorator.run(DefaultThreadFactory.java:137)
	at java.lang.Thread.run(Thread.java:745)

$`)
	parsed, threaddump := parser.ParseThreaddump()
	assert := assert.New(t)
	assert.True(parsed == "", "ParseThread should succeed on valid input")
	assert.Equal(1, len(threaddump.Threads))
	assert.Equal("Full thread dump Java HotSpot(TM) 64-Bit Server VM (25.20-b23 mixed mode):", threaddump.Header)
	assert.Equal(byte('$'), parser.Current(), "ParseThread must consume the entire thread")
}
