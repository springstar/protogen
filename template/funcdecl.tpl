<%= package%>
<%= for (p) in imports() { %>
<%= "import "%> "<%=p%>"
<% } %>

<%="var descriptors map[int32]*desc.MessageDescriptor = make(map[int32]*desc.MessageDescriptor)"%>

<%= for (n) in names { %>
<%="func parse"%><%=n%>(<%="id int32, bytes []byte"%>) *<%="pb."%><%=n%> {
    <%= "dmsg := dynamic.NewMessage(md)"%>
}
<% } %>

