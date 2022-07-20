$.get('path/template.mustache', function (partialTemplate) {

	var html = Mustache.render(template, obj, {
	  templateRef: partialTemplate
	});
  
	// ...
  
  }, 'text');
  

  cmd.get('header.html.mustache', function (partialHead) {
	var header = Mustache.render(template, obj, {
		templateHead: partialHeader
	});
}, 'text');