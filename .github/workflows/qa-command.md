QA :{{.status}}_circle:

- :{.build.status}_circle: [build]
- :{.deploy.status}_circle: [deploy]
- :{.verify.status}_circle: [verify]
- :{.cleanup.status}_circle: [cleanup]

[build]: {{.build.link}}
[deploy]: {{.deploy.link}}
[verify]: {{.verify.link}}
[cleanup]: {{.cleanup.link}}

<!--
Expected Payload:
{
    "status": "red,green,yellow,white",
    "build": {
        "status": "red,green,yellow,white",
        "link": "somewhere"
    },
    "deploy": ...
    "verify": ...
    "cleanup": ...
}
-->
