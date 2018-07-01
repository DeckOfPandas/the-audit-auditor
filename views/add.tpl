<html>
<head>
	<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">
	<link rel="stylesheet" href="//cdn.rawgit.com/necolas/normalize.css/master/normalize.css">
	<link rel="stylesheet" href="//cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css">	
	<style>
		body {
			padding: 100px;
			color: black;
			font-weight: bold;
		}
		.field{
			position: relative;
			width: 200px;
		}
	</style>
	<script
	  src="https://code.jquery.com/jquery-2.2.4.min.js"
	  integrity="sha256-BbhdlvQf/xTY9gja0Dq3HiwQF8LaCRTXxZKRutelT44="
	  crossorigin="anonymous"></script>
</head>
<body>
	<form action="" method="post">
		{{.AuditForm | renderform}}
		<button class="field" type="submit">Submit!</button>
	</form>
	<script>
	$('.field').on('mouseenter',function(e){
		var maxX = $(window).width() - $(this).width();
		var maxY = $(window).height() - $(this).height();    
		$(this).css({
			'left':getRandomInt(0, maxX/4),
			'top':getRandomInt(0, maxY/4)
		});
	});
	function getRandomInt(min, max) {
		return Math.floor(Math.random() * (max - min + 1)) + min;
	}	
	</script>
</body>
</html>
