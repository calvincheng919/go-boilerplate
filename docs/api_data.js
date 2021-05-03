define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./docs/main.js",
    "group": "C:\\Users\\JordanCotter\\Development\\lemonbrew\\go-service-boilerplate\\docs\\main.js",
    "groupTitle": "C:\\Users\\JordanCotter\\Development\\lemonbrew\\go-service-boilerplate\\docs\\main.js",
    "name": ""
  },
  {
    "type": "get",
    "url": "/version",
    "title": "",
    "name": "Version",
    "description": "<p>displays version</p>",
    "group": "misc",
    "version": "0.0.0",
    "examples": [
      {
        "title": "Curl Example:",
        "content": "curl http://localhost:8080/version",
        "type": "curl"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "version",
            "description": "<p>current version</p>"
          }
        ]
      }
    },
    "filename": "./src/main.go",
    "groupTitle": "misc"
  }
] });
