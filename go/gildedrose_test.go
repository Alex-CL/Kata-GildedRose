package gildedrose_test

import (
	"testing"

	"github.com/alex-cl/kata-gildedrose/go"
)

func Test_Normal_Item(t *testing.T) {
	iq := 20
	var items = []*gildedrose.Item{
		{"Fresh Water", 10, iq},
		{"Apples", 10, 0},
		{"Pears", 0, iq},
	}

	gildedrose.UpdateQuality(items)

	expectedQualities := []int{
		iq - 1,
		0,
		iq - 2,
	}
	
	for i := range expectedQualities {
		if items[i].Quality != expectedQualities[i] {
			t.Errorf("Quality: Expected %d but got %d", expectedQualities[i], items[i].Quality)
		}
	}
}

func Test_Aged_Brie(t *testing.T) {
	iq := 20
	var items = []*gildedrose.Item{
		{gildedrose.AgedBrie, 5, iq},
		{gildedrose.AgedBrie, 0, iq},
		{gildedrose.AgedBrie, 0, gildedrose.MaxQuality},
	}

	gildedrose.UpdateQuality(items)
	
	expected := []gildedrose.Item{
		{gildedrose.AgedBrie, 4, iq + 1},
		{gildedrose.AgedBrie, -1, iq + 2},
		{gildedrose.AgedBrie, -1, gildedrose.MaxQuality},
	}
	
	for i := range items {
		if items[i].Quality != expected[i].Quality {
			t.Errorf("Quality: Expected %d but got %d", expected[i].Quality, items[i].Quality)
		}
		if items[i].SellIn != expected[i].SellIn {
			t.Errorf("SellIn: Expected %d but got %d", expected[i].SellIn, items[i].SellIn)
		}
	}
}

func Test_Sulfuras(t *testing.T) {
	iq := 80
	isn := 10
	var items = []*gildedrose.Item{
		{gildedrose.Sulfuras, isn, iq},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality != iq {
		t.Errorf("Quality must not change for Sulfuras: Expected %d but got %d", iq, items[0].Quality)
	}
	
	if items[0].SellIn != isn {
		t.Errorf("SellIn must not change for Sulfuras: Expected %d but got %d", isn, items[0].SellIn)
	}
}

func Test_Backstage(t *testing.T) {
	iq := 30
	var items = []*gildedrose.Item{
		{gildedrose.Backstage, 40, iq},
		{gildedrose.Backstage, 10, iq},
		{gildedrose.Backstage, 5, iq},
		{gildedrose.Backstage, 0, iq},
	}

	gildedrose.UpdateQuality(items)

	expectedQualities := []int{
		iq + 1,
		iq + 2,
		iq + 3,
		0,
	}

	for i := range expectedQualities {
		if items[i].Quality != expectedQualities[i] {
			t.Errorf("Quality: Expected %d but got %d", expectedQualities[i], items[i].Quality)
		}
	}
}

func Test_Conjured(t *testing.T) {
	iq := 30
	var items = []*gildedrose.Item{
		{gildedrose.ConjuredPrefix + " wand", 40, iq},
		{gildedrose.ConjuredPrefix + " bread", 10, iq},
		{gildedrose.ConjuredPrefix + " beer", 0, 0},
	}
	
	gildedrose.UpdateQuality(items)
	
	expectedQualities := []int{
		iq - 2,
		iq - 2,
		0,
	}

	for i := range expectedQualities {
		if items[i].Quality != expectedQualities[i] {
			t.Errorf("Quality: Expected %d but got %d", expectedQualities[i], items[i].Quality)
		}
	}
	
}

