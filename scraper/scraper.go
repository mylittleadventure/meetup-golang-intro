package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const perPage = 25
const bookingURL = "https://www.booking.com/searchresults.en-gb.html?aid=807124&sid=53bce7046c98b8fcc190ba4c092250bc&sb=1&src=country&src_elem=sb&error_url=https%3A%2F%2Fwww.booking.com%2Fcountry%2Fus.en-gb.html%3Faid%3D807124%3Bsid%3D53bce7046c98b8fcc190ba4c092250bc%3Binac%3D0%26%3B&ss=Nice%2C+Provence-Alpes-Côte+d'Azur%2C+France&is_ski_area=&checkin_monthday=10&checkin_month=2&checkin_year=2019&checkout_monthday=11&checkout_month=2&checkout_year=2019&no_rooms=1&group_adults=2&group_children=0&from_sf=1&ss_raw=nice&ac_position=0&ac_langcode=en&ac_click_type=b&dest_id=-1454990&dest_type=city&iata=NCE&place_id_lat=43.69808&place_id_lon=7.269624&search_pageview_id=00f8759fa84200ab&search_selected=true&search_pageview_id=00f8759fa84200ab&ac_suggestion_list_length=5&ac_suggestion_theme_list_length=0"

// Hotel ...
type Hotel struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	RatingAvg string `json:"rating_avg"`
}

// ScrapeBooking launches scrapping of booking
func ScrapeBooking(client *http.Client, pages int) (h []Hotel) {
	c := make(chan []Hotel, pages)

	for i := 0; i < pages; i++ {
		go Scrape(client, bookingURL+fmt.Sprintf("&offset=%d&rows=%d", pages*perPage, perPage), c)
	}

	for i := 0; i < pages; i++ {
		h = append(h, <-c...)
	}

	return
}

// Scrape scraps a url & passes the hotels found in channel
func Scrape(client *http.Client, url string, c chan []Hotel) {
	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	hotels := []Hotel{}
	// Find all links and process them with the function defined earlier
	document.Find(".sr_item").Each(func(index int, element *goquery.Selection) {
		hotels = append(hotels, Hotel{
			Title:     strings.TrimSpace(element.Find(".sr-hotel__name").Text()),
			Thumbnail: element.Find(".hotel_image").AttrOr("src", ""),
			RatingAvg: strings.TrimSpace(element.Find(".bui-review-score__badge").Text()),
		})
	})

	c <- hotels
}
