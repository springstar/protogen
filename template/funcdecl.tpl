<%= package%>
<%= for (p) in imports() { %>
<%= "import "%> "<%=p%>"
<% } %>

<%="var descriptors map[int32]*desc.MessageDescriptor = make(map[int32]*desc.MessageDescriptor)"%>

<%= for (n) in names { %>
<%= decl(n) <%= lbrack() %>
	<%= "dmsg := dynamic.NewMessage(md)"%>
<%= rbrack() %>
<% } %>
