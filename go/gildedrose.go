package gildedrose

import (
	"strings"
)

const (
	MaxQuality = 50
	
	AgedBrie = "Aged Brie"
	Backstage = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras = "Sulfuras, Hand of Ragnaros"
	ConjuredPrefix = "Conjured"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := range items {
		if items[i].Name == Sulfuras {
			continue
		}

		items[i].SellIn--
	
		if strings.HasPrefix(items[i].Name, ConjuredPrefix) {
			if items[i].Quality >= 2 {
				items[i].Quality -= 2
			} else {
				items[i].Quality = 0
			}
			
			continue
		}
		
		if items[i].Name == AgedBrie {
			items[i].Quality++

			if items[i].SellIn < 0 {		
				items[i].Quality++
			}
			
			if items[i].Quality > MaxQuality {
				items[i].Quality = MaxQuality
			}

			continue
		}
		
		if items[i].Name == Backstage {
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
