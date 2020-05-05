package main

import (
	"bytes"
	"html/template"
	"io"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("https://chroju.github.io/atom.xml")
	ghitems := feed.Items[:3]
	feed, _ = fp.ParseURL("https://chroju.hatenablog.jp/feed")
	hbitems := feed.Items[:3]

	tmpl := template.Must(template.New("index.html").Parse(htmlTemplate))
	buf := new(bytes.Buffer)
	w := io.Writer(buf)

	err := tmpl.ExecuteTemplate(w, "base", struct {
		GitHubIOEntries   []*gofeed.Item
		HatenaBlogEntries []*gofeed.Item
	}{
		GitHubIOEntries:   ghitems,
		HatenaBlogEntries: hbitems,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(buf.Bytes()),
	}, nil
}

func main() {
	lambda.Start(handler)
}

const htmlTemplate = `
{{define "base"}}
<!doctype html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<link rel="icon" href="https://secure.gravatar.com/avatar/542bf1e833425f6ab7bc7bd7238a4792?s=24" />
	<script src="https://kit.fontawesome.com/215264fa68.js"></script>
	<title>chroju</title>
</head>
<body>
	<h1>chroju</h1>
	<img src="https://secure.gravatar.com/avatar/542bf1e833425f6ab7bc7bd7238a4792?s=250" alt="chroju" />
	<a href="https://github.com/chroju"><i class="fab fa-github-alt fa-2x"></i></a>
	<a href="https://twitter.com/chroju"><i class="fab fa-twitter fa-2x"></i></a>
	<a href="https://www.instagrafa-instagramm.com/chroju"><i class="fab fa-instagram fa-2x"></i></a>
	<a href="https://speakerdeck.com/chroju"><i class="fab fa-speaker-deck fa-2x"></i></a>

    <h2>Who</h2>
    <dl>
        <dt>Job</dt>
            <dd>Site Reliability Engineer</dd>
        <dt>Location</dt>
            <dd>Tokyo, Japan</dd>
        <dt>Skills</dt>
            <dd>Automation / Practical Monitoring / Infrastructure as Code / Documentation</dd>
        <dt>Techs</dt>
            <dd>AWS / Terraform / Kubernetes / VMware / Go / Python / bash ... etc</dd>
        <dt>Contact</dt>
            <dd>chor.chroju at gmail.com</dd>
    </dl>

    <h2>Experience</h2>
    <dl>
        <dt>GLOBIS Corporation</dt>
            <dd>Site Reliability Engineer</dd>
            <dd>Apr 2020 - current</dd>
        <dt>Freelancer</dt>
            <dd>Site Reliability Engineer</dd>
            <dd>Jun 2019 - Mar 2020</dd>
        <dt>Quants Research Inc.</dt>
            <dd>Web Operation Engineer</dd>
            <dd>Jun 2015 - May 2019</dd>
        <dt>TIS Inc.</dt>
            <dd>System Engineer</dd>
            <dd>Apr 2011 - May 2015</dd>
    </dl>

    <h2>Education</h2>
    <dl>
        <dt>Teikyo University (Distance Learning)</dt>
            <dd>Bachelor of Science in Information Technology</dd>
            <dd>Apr 2019 - current</dd>
        <dt>Hitotsubashi University</dt>
            <dd>Bachelor of Social Science</dd>
            <dd>Apr 2007 - Mar 2011</dd>
    </dl>

    <h2>Blogs (recent entries)</h2>
	<h3><a href="https://chroju.github.io/">the world as code</a></h3>
	<p>about technologies</p>
    <ul>
        {{range $entry := .GitHubIOEntries }}
        <li><a href="{{$entry.Link}}">{{$entry.Title}}</a></li>
        {{end}}
    </ul>

    <h3><a href="https://chroju.hatenablog.jp/">the world was not enough</a></h3>
	<p>about cultures</p>
    <ul>
        {{range $entry := .HatenaBlogEntries }}
        <li><a href="{{$entry.Link}}">{{$entry.Title}}</a></li>
        {{end}}
	</ul>
	
	<footer>
		<p>This page is generated by <a href="https://www.netlify.com/products/functions/">Netlify Functions</a>.</p>
	</footer>

</body>
{{end}}
`
