package gildedrose

import (
	"strings"
)

const (
	maxQuality = 50
	
	agedBrie = "Aged Brie"
	backstage = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras = "Sulfuras, Hand of Ragnaros"
	conjuredPrefix = "Conjured"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := range items {
		if items[i].Name == sulfuras {
			continue
		}

		items[i].SellIn--
	
		if strings.HasPrefix(items[i].Name, conjuredPrefix) {
			if items[i].Quality >= 2 {
				items[i].Quality -= 2
			} else {
				items[i].Quality = 0
			}
			
			continue
		}
		
		if items[i].Name == agedBrie {
			items[i].Quality++

			if items[i].SellIn < 0 {		
				items[i].Quality++
			}
			
			if items[i].Quality > maxQuality {
				items[i].Quality = maxQuality
			}

			continue
		}
		
		if items[i].Name == backstage {
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
		
		items[i].Quality--
		
		if items[i].SellIn < 0 {
			items[i].Quality--
		}
		
		if items[i].Quality < 0 {		
			items[i].Quality = 0
		}
	}
}
