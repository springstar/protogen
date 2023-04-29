<%= package%>

import  (
    "github.com/springstar/protogen/pb"
    "github.com/jhump/protoreflect/desc"
    "github.com/jhump/protoreflect/dynamic"   
)

<%="var descriptors map[int32]*desc.MessageDescriptor = make(map[int32]*desc.MessageDescriptor)"%>

func AddDescriptor(id int32, desc *desc.MessageDescriptor) {
    descriptors[id] = desc
}

<%= for (n) in names { %>
<%="func parse"%><%=n%>(<%="id int32, bytes []byte"%>) *<%="pb."%><%=n%> {
    <%= "msg := "%> &<%="pb."%><%=n%>{}
    md := descriptors[id]
    dmsg := dynamic.NewMessage(md)
    dmsg.Unmarshal(bytes)
    dmsg.ConvertTo(msg)
    return msg
}
<% } %>


<%= for (n) in names { %>
<%="func serialize"%><%=n%>(<%=params(n)%>) *<%="pb."%><%=n%> {
    msg := &pb.<%=n%>{

    }
    return msg
}
<% } %>