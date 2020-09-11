<!-- forms.html -->

<html>
<head>
	<title>User info</title>
</head>
<body>
    <h1>Thanks for your message!</h1>
    <li>Btc Addr{{.Balance}}</li>    
    <li>Btc Addr{{.Balance}}</li>
    <li>Btc Addr{{.Balance}}</li>


     <h1>Check BALANCE </h1>
        <form action="/show" method="POST">
        <label>Address:</label><br />
        <input type="text" name="address"><br />
        <input type="submit">
    </form>

</body>
</html>