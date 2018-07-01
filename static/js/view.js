$(document).ready(function(){
	$('#table').DataTable({
		"processing": true,
		"serverSide": true,
		"ajax": '/datatables/data/json',
		"aLengthMenu": [[5,15, 25, 50, 100, 200, 500], ['5','15', '25', '50', '100', '200', '500']],
		"iDisplayLength": 15,
		"columns": [
			{ "data": "Id" },
			{ "data": "Title" },
			{ "data": "Summary" },
			{ "data": "Author" },
			{ "data": "Email" },
			{ "data": "Location" },
			{ "data": "Days" }
		]
	})
})
