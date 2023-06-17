package notify

import "github.com/opencamp-hq/core/models"

type EmailData struct {
	Campground     *models.Campground
	StartDate      string
	EndDate        string
	AvailableSites models.Campsites
}

var EmailTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Good news! {{.Campground.Name}} is available</title>
	<style>
		body { font-family: Arial, sans-serif; }
		h1 { color: #333; }
		p { color: #666; }
		table {
			width: 100%;
			border-collapse: collapse;
		}
		th, td {
			border: 1px solid #dddddd;
			text-align: left;
			padding: 8px;
		}
		tr:nth-child(even) {
			background-color: #f2f2f2;
		}
		#footer {
			margin-top: 50px;
			color: #999;
			font-size: 0.8em;
		}
	</style>
</head>

<body>
	<h1>{{.Campground.Name}} now has campsites available!</h1>
	<p>{{.Campground.ParentName}}, near {{.Campground.City}}</p>
	<p><strong>Check-in:</strong> {{.StartDate}}</p>
	<p><strong>Check-out:</strong> {{.EndDate}}</p>

	<table>
		<tr>
			<th>Campsite</th>
			<th>Loop</th>
			<th>Type</th>
			<th>Use</th>
			<th>Min People</th>
			<th>Max People</th>
			<th>Book Now</th>
		</tr>
		{{range .AvailableSites}}
		<tr>
			<td>{{.Site}}</td>
			<td>{{.Loop}}</td>
			<td>{{.CampsiteType}}</td>
			<td>{{.TypeOfUse}}</td>
			<td>{{.MinNumPeople}}</td>
			<td>{{.MaxNumPeople}}</td>
			<td><a href="https://www.recreation.gov/camping/campsites/{{.CampsiteID}}">Book Now</a></td>
		</tr>
		{{end}}
	</table>
	<div id="footer">
		<hr>
		<p style="font-size: small;"><em>Brought to you by OpenCamp</em></p>
	</div>
</body>
</html>
`
