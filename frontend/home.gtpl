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
                <form action="{{if .Post.ID}}/{{.Post.ID}}/update{{ else}}/create{{end}}"
                      method="{{if .Post.ID}}PUT{{else}}POST{{end}}">
                    <h4>{{if .Post.ID}}Update Blog{{else}}Create New Blog{{end}}</h4>
                    <div class="mb-3">
                        <label for="title" class="form-label">Blog Title</label>
                        <input type="text" name="title" class="form-control" id="title"
                               value="{{if .Post.ID}}{{.Post.Title}}{{end}}"
                               placeholder="Go is the best">
                    </div>
                    <div class="mb-3">
                        <label for="details" class="form-label">Blog Details</label>
                        <textarea class="form-control" id="details"
                                  placeholder="Enter Details..."
                                  name="details" rows="3">{{if .Post.ID}}{{.Post.Details}}{{end}}</textarea>
                    </div>
                    <div class="mb-3">
                        <input type="hidden" name="blogId" value="{{if .Post.ID}}{{.Post.ID}}{{end}}">
                            <button type="submit" class="btn btn-dark btn-block float-right form-control">
                                {{if .Post.ID}}Update Blog {{else}}Create New Blog Post{{end}}
                            </button>
                    </div>
                </form>
            </div>
            <div class="col-6">
                <h4>Previous Blog Post</h4>
                {{range .Data.Blogs}}

                <div class="row mb-3">
                    <a href="#">{{ .Title}}</a>
                    <div class="row">
                        <a href="/{{.ID}}/edit" class="col-2 btn btn-primary btn-sm">Edit</a>
                        <a href="/{{.ID}}/delete" onclick="return confirm('Are you sure you want to delete?')"
                           class="col-2 btn btn-danger btn-sm">Delete</a>
                    </div>
                </div>
                {{end}}
            </div>
        </div>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
            crossorigin="anonymous"></script>
</body>
</html>