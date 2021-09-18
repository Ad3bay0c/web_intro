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
            <div class="col-6">
                {{if .Message}}
                <div class="alert alert-{{.Color}}">
                    {{ .Message}}
                </div>
                {{end}}
                <form action="/create" method="POST">
                    <div class="mb-3">
                        <label for="title" class="form-label">Blog Title</label>
                        <input type="text" name="title" class="form-control" id="title"
                               placeholder="Go is the best">
                    </div>
                    <div class="mb-3">
                        <label for="details" class="form-label">Blog Details</label>
                        <textarea class="form-control" id="details"
                                  placeholder="Enter Details..."
                                  name="details" rows="3"></textarea>
                    </div>
                    <div class="mb-3">
                            <button type="submit" class="btn btn-dark btn-block float-right form-control">Create New Blog Post</button>
                    </div>
                </form>
            </div>
            <div class="col-6"></div>
        </div>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
            crossorigin="anonymous"></script>
</body>
</html>