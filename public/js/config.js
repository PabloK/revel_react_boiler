requirejs.config({
  waitSeconds: 1,
	enforceDefine: true,
	"baseUrl": "public/js/",
	paths: {
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
		flux: [	
    //TODO: Find CDN for flux
		'lib/flux-2.0.2'
		],
		cookies: [
		'lib/cookie-1.2.1.min'
		],
		log: [
		'log'
		]
	}
});
require(["main"]);
