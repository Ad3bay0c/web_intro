<html>
<head>
    <title>Blog</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU"
          crossorigin="anonymous">
</head>
<body>
<header class="mb-5">
    {{template "navbar.gtpl"}}
</header>
<div class="container">
    <div class="row justify-content-center">
        {{with .Blogs}}{{range $idx, $blog := . }}
        <div class="col-12 mb-3">
            <div class="card">
                <div class="card-body">
                    <a href="#" class="card-title text-success" style="text-decoration-line: none">{{ $blog.Title}}</a>
                    <div class="row mt-3 card-subtitle ">
                        <div class="text-muted col-3">
                            <h6>Comments: <b class="badge bg-success">12</b></h6>
                        </div>
                        <div class="text-muted col-3">
                            <h6>Views: <b class="badge bg-success">12</b></h6>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        {{end}}{{end}}
    </div>
</div>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
        crossorigin="anonymous"></script>
</body>
</html>