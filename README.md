[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# OpenCamp

Simple Go library for interacting with the recreation.gov API.

## Usage

### Check campground availability
```
l := log15.New()
c := client.New(l, 10*time.Second)

sites, err := c.Availability(campgroundID, start, end)
if err != nil {
  // handle err
}

if len(sites) == 0 {
  fmt.Println("Sorry we didn't find any available campsites!")
} else {
  fmt.Println("The following sites are available for those dates:")
  for _, s := range sites {
    fmt.Printf(" - Site %-15s Book at: https://www.recreation.gov/camping/campsites/%s\n", s.Site, s.CampsiteID)
  }
}
```

### Poll campground availability
```
l := log15.New()
c := client.New(l, 10*time.Second)
ctx := context.Background()

// Blocking operation.
sites, err := c.Poll(ctx, campgroundID, start, end, interval)
if err != nil {
  // handle err
}

fmt.Println("The following sites are available for those dates:")
for _, s := range sites {
  fmt.Printf(" - Site %-15s Book at: https://www.recreation.gov/camping/campsites/%s\n", s.Site, s.CampsiteID)
}
```

### Search for a campground Id
```
l := log15.New()
c := client.New(l, 10*time.Second)
campgrounds, err := c.Suggest(args[0])
if err != nil {
  // handle err
}
if len(campgrounds) == 0 {
  fmt.Println("Sorry, no campgrounds with that name were found.")
  return
}

bytes, err := json.MarshalIndent(campgrounds, "", "  ")
if err != nil {
  // handle err
}

fmt.Println(string(bytes))
```

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/opencamp-hq/core?style=for-the-badge
[contributors-url]: https://github.com/opencamp-hq/core/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/opencamp-hq/core?style=for-the-badge
[forks-url]: https://github.com/opencamp-hq/core/network/members
[stars-shield]: https://img.shields.io/github/stars/opencamp-hq/core?style=for-the-badge
[stars-url]: https://github.com/opencamp-hq/core/stargazers
[issues-shield]: https://img.shields.io/github/issues/opencamp-hq/core?style=for-the-badge
[issues-url]: https://github.com/opencamp-hq/core/issues
[license-shield]: https://img.shields.io/github/license/opencamp-hq/core?style=for-the-badge
[license-url]: https://github.com/opencamp-hq/core/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/kylechadha
