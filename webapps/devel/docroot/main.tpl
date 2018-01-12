<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="/static/libs/w3css/4/w3mobile.css">
	<link rel="stylesheet" href="/static/libs/w3-theme-dark-grey.css">
	<!-- <title>FIXME</title> -->
</head>
<body class="w3-theme-dark">

	<!-- HEADER -->
	<header class="w3-container w3-bar w3-theme">
		<h1 class="w3-bar-item">JCMS Devel</h1>
	</header>

	<!-- MENU -->
	<nav class="w3-container w3-bar">
		<a class="w3-bar-item w3-button" href="/">Home</a>
	</nav>

	<!-- MAIN -->
	<div class="w3-container w3-theme-light">
		{{block "main" .}}{{end}}
	</div>

	<!-- FOOTER -->
	<footer class="w3-container w3-margin-top w3-small w3-theme-dark">
		<p>jcms v0.0</p>
	</footer>
</body>
</html>
