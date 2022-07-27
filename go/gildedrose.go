package gildedrose

import (
	"strings"
)

const (
	MAX_QUALITY = 50
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := range items {
		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}
	
		if strings.HasPrefix(items[i].Name, "Conjured") {
			items[i].SellIn--
			
			if items[i].Quality >= 2 {
				items[i].Quality -= 2
			} else {
				items[i].Quality = 0
			}
			
			continue
		}
		
		if items[i].Name == "Aged Brie" {
			items[i].SellIn--
			items[i].Quality++

			if items[i].SellIn < 0 {		
				items[i].Quality++
			}
			
			if items[i].Quality > MAX_QUALITY {
				items[i].Quality = MAX_QUALITY
			}

			continue
		}
		
		if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			items[i].SellIn--
			
			if items[i].SellIn < 0 {
				items[i].Quality = 0
			} else if items[i].SellIn <= 5 {
				items[i].Quality += 3
			} else if items[i].SellIn <= 10 {
				items[i].Quality += 2
			} else {
				items[i].Quality++
			}			
			
			continue
		}
		
		items[i].SellIn--
		items[i].Quality--
		
		if items[i].SellIn < 0 {
			items[i].Quality--
		}
		
		if items[i].Quality < 0 {		
			items[i].Quality = 0
		}
	}
}
