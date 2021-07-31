# vnet
virtual net interface

golang provides net.Listener and net.Conn interfaces to wrap the network. A lot of interesting tcp codes are built on these two interfaces. So as long as these two interfaces are met, these codes can be used outside of tcp, this is what vnet does.

# PipeListener

The ListenPipe function returns a PipeListener. PipeListener implements the net.Listener interface in memory. Use PipeListener.Dial or PipeListener.DialContext to dial.

One of the most typical applications is to use PipeListener to provide grpc and grpc-gateway services in the program. The grpc and grpc-gateway use memory copy communication instead of socket communication.

# reverse.Dialer reverse.Listener

reverse.Dialer reverse.Listener used to reverse the server client. This allows socket.connect to work in the server role and socket.accept to work in the client role.

The most typical use is that the server role works on the internal network, and the client role works on the external network. For example, reverse Trojan horses or programs that need to break through the intranet to provide services. Imagine how easy it is to write a reverse Trojan using grpc.