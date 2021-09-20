<html>
<head>
    <title>Blog</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU"
          crossorigin="anonymous">
</head>
<body>
<header class="mb-5">
    {{template "Navbar"}}
</header>
<div class="container">
    <div class="row">
        <div class="col-md-12 float-right mb-2">
            <a href="/" class="btn btn-dark">Back</a>
        </div>
    </div>
    <div class="shadow p-3 mb-5 bg-body rounded">
        <h4>
            <strong>{{.Post.Title}}</strong>
        </h4>
        <hr>
        <p>{{.Post.Details}}</p>
    </div>

    <div class="row">
        {{if .Message}}
        <div class="alert alert-{{.Color}}">
            {{ .Message}}
        </div>
        {{end}}
        <div class="col-md-12 col-sm-12 col-lg-4">
            <form action="/{{.Post.ID}}/comment/create" method="post">
                <textarea class="form-control" name="comment" rows="5" placeholder="Write your comment" required></textarea>
                <br>
                <button class="btn btn-dark" type="submit">Comment</button>
            </form>
        </div>
        <div class="col-md-12 col-sm-12 col-lg-6">
            <h5>Comments <b class="badge bg-success">{{.Post.Comment}}</b></h5>
            <hr>
            {{with .Post.Comments}}{{range $key, $comment := .}}
                <div class="mb-3" style="border-bottom: dotted lightgreen">{{$comment.Comment}}</div>
            {{end}}{{end}}
        </div>
    </div>

</div>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
        crossorigin="anonymous"></script>
</body>
</html>