requirejs.config({
	enforceDefine: true,
	"baseUrl": "public/js/",
	paths: {
		bootstrap: [
		'//maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min',
		//If the CDN location fails, load from this location
		'lib/bootstrap-3.3.4.min'
		],
		jquery: [
		'//ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min',
		//If the CDN location fails, load from this location
		'lib/jquery-2.1.3.min'
		],
		react: [	
		'//cdnjs.cloudflare.com/ajax/libs/react/0.13.1/react.min',
		//If the CDN location fails, load from this location
		'lib/react-0.13.1.min'
		],
		cookies: [
		'lib/cookie-1.2.1.min'
		]
	}
});
requirejs(["main"]);
