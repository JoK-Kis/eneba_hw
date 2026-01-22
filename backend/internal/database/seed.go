package database

import (
	"log"

	"github.com/JoK-Kis/eneba-backend/internal/models"
)

func Seed() {
	var platformCount int64
	DB.Model(&models.Platform{}).Count(&platformCount)
	if platformCount > 0 {
		log.Println("Database already seeded, skipping seed")
		return
	}

	platforms := []models.Platform{
		{Name: "Nintendo", IconURL: "./assets/platform/nintendo.svg"},
		{Name: "Steam", IconURL: "./assets/platform/steam.svg"},
		{Name: "Xbox", IconURL: "./assets/platform/xbox.svg"},
		{Name: "PlayStation", IconURL: "./assets/platform/psn.svg"},
		{Name: "EA", IconURL: "./assets/platform/EA.svg"},
		{Name: "Origin", IconURL: "./assets/platform/origin.svg"},
		{Name: "Rockstar", IconURL: "./assets/platform/rockstar_social_club.svg"},
	}

	for _, platform := range platforms {
		if err := DB.Create(&platform).Error; err != nil {
			log.Printf("Error creating platform %s: %v", platform.Name, err)
		}
	}

	var nintendoPlatform models.Platform
	DB.Where("name = ?", "Nintendo").First(&nintendoPlatform)

	var steamPlatform models.Platform
	DB.Where("name = ?", "Steam").First(&steamPlatform)

	var xboxPlatform models.Platform
	DB.Where("name = ?", "Xbox").First(&xboxPlatform)

	var rockstarPlatform models.Platform
	DB.Where("name = ?", "Rockstar").First(&rockstarPlatform)

	var eaPlatform models.Platform
	DB.Where("name = ?", "EA").First(&eaPlatform)

	var playstationPlatform models.Platform
	DB.Where("name = ?", "PlayStation").First(&playstationPlatform)

	var originPlatform models.Platform
	DB.Where("name = ?", "Origin").First(&originPlatform)

	var greenGiftPlatform models.Platform
	DB.Where("name = ?", "Green Gift").First(&greenGiftPlatform)

	products := []models.Product{
		{
			Name:            "Split Fiction (Xbox Series X|S) XBOX LIVE Key EUROPE",
			ImageURL:        "./assets/product/split.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "EU",
			OriginalPrice:   59.99,
			CurrentPrice:    35.99,
			DiscountPercent: 40,
			CashbackAmount:  7.00,
			Likes:           1789,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2: Ultimate Edition (PC) Rockstar Games Launcher Key EUROPE",
			ImageURL:        "./assets/product/rdSpecial.jpg",
			PlatformID:      rockstarPlatform.ID,
			Region:          "EU",
			OriginalPrice:   69.99,
			CurrentPrice:    52.49,
			DiscountPercent: 25,
			CashbackAmount:  0.00,
			Likes:           1542,
			InStock:         true,
		},
		{
			Name:            "Split Fiction EA App Key (PC) GLOBAL",
			ImageURL:        "./assets/product/split.jpg",
			PlatformID:      eaPlatform.ID,
			Region:          "GLOBAL",
			OriginalPrice:   49.99,
			CurrentPrice:    44.99,
			DiscountPercent: 10,
			CashbackAmount:  5.00,
			Likes:           2341,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2 (PC) Rockstar Games Launcher Key EUROPE",
			ImageURL:        "./assets/product/rd.jpg",
			PlatformID:      rockstarPlatform.ID,
			Region:          "EU",
			OriginalPrice:   59.99,
			CurrentPrice:    47.99,
			DiscountPercent: 20,
			CashbackAmount:  8.00,
			Likes:           876,
			InStock:         true,
		},
		{
			Name:            "Split Fiction (Xbox Series X|S) XBOX LIVE Key GLOBAL",
			ImageURL:        "./assets/product/split.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "GLOBAL",
			OriginalPrice:   54.99,
			CurrentPrice:    38.49,
			DiscountPercent: 30,
			CashbackAmount:  12.00,
			Likes:           654,
			InStock:         true,
		},
		{
			Name:            "Split Fiction (Nintendo Switch 2) eShop Key EUROPE",
			ImageURL:        "./assets/product/splitNintendo.jpg",
			PlatformID:      nintendoPlatform.ID,
			Region:          "EU",
			OriginalPrice:   59.99,
			CurrentPrice:    59.99,
			DiscountPercent: 0,
			CashbackAmount:  0.00,
			Likes:           432,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2 (PC) Green Gift Key EUROPE",
			ImageURL:        "./assets/product/rdbundle.jpg",
			PlatformID:      greenGiftPlatform.ID,
			Region:          "EU",
			OriginalPrice:   64.99,
			CurrentPrice:    45.49,
			DiscountPercent: 30,
			CashbackAmount:  15.00,
			Likes:           1876,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2: Ultimate Edition Rockstar Games Launcher Key UNITED STATES/EMEA",
			ImageURL:        "./assets/product/rdbundle.jpg",
			PlatformID:      rockstarPlatform.ID,
			Region:          "US/EMEA",
			OriginalPrice:   49.99,
			CurrentPrice:    42.49,
			DiscountPercent: 15,
			CashbackAmount:  7.00,
			Likes:           521,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2: Ultimate Edition Rockstar Games Launcher Key UNITED STATES/EUROPE",
			ImageURL:        "./assets/product/rdSpecial.jpg",
			PlatformID:      rockstarPlatform.ID,
			Region:          "US/EUROPE",
			OriginalPrice:   59.99,
			CurrentPrice:    35.99,
			DiscountPercent: 40,
			CashbackAmount:  20.00,
			Likes:           789,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2 (Xbox One) Xbox Live Key EUROPE",
			ImageURL:        "./assets/product/rdXbox.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "EU",
			OriginalPrice:   69.99,
			CurrentPrice:    49.99,
			DiscountPercent: 29,
			CashbackAmount:  8.00,
			Likes:           1234,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2 - Ultimate Edition (Xbox One) Xbox Live Key EUROPE",
			ImageURL:        "./assets/product/rdbundle.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "EU",
			OriginalPrice:   69.99,
			CurrentPrice:    52.99,
			DiscountPercent: 24,
			CashbackAmount:  10.00,
			Likes:           987,
			InStock:         true,
		},
		{
			Name:            "Red Dead Redemption 2 Rockstar Games Launcher Key EMEA",
			ImageURL:        "./assets/product/rd.jpg",
			PlatformID:      rockstarPlatform.ID,
			Region:          "EMEA",
			OriginalPrice:   59.99,
			CurrentPrice:    29.99,
			DiscountPercent: 50,
			CashbackAmount:  5.00,
			Likes:           3456,
			InStock:         true,
		},
		{
			Name:            "EA SPORTS™ FIFA 23 Standard Edition Xbox Series X|S Key GLOBAL",
			ImageURL:        "./assets/product/fifa.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "GLOBAL",
			OriginalPrice:   79.99,
			CurrentPrice:    39.99,
			DiscountPercent: 50,
			CashbackAmount:  8.00,
			Likes:           2100,
			InStock:         true,
		},
		{
			Name:            "EA SPORTS™ FIFA 23 Standard Edition Xbox One Key GLOBAL",
			ImageURL:        "./assets/product/fifaXbox.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "GLOBAL",
			OriginalPrice:   59.99,
			CurrentPrice:    35.99,
			DiscountPercent: 40,
			CashbackAmount:  7.00,
			Likes:           1789,
			InStock:         true,
		},
		{
			Name:            "EA SPORTS™ FIFA 23 Standard Edition Xbox Series X|S Key EUROPE",
			ImageURL:        "./assets/product/fifaXbox.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "EU",
			OriginalPrice:   39.99,
			CurrentPrice:    24.99,
			DiscountPercent: 38,
			CashbackAmount:  3.00,
			Likes:           456,
			InStock:         true,
		},
		{
			Name:            "EA SPORTS™ FIFA 23 Ultimate Edition Xbox One & Xbox Series X|S Key EUROPE",
			ImageURL:        "./assets/product/fifaXbox2.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "EU",
			OriginalPrice:   39.99,
			CurrentPrice:    24.99,
			DiscountPercent: 38,
			CashbackAmount:  3.00,
			Likes:           456,
			InStock:         true,
		},
		{
			Name:            "EA SPORTS™ FIFA 23 Standard Edition Xbox One Key EUROPE",
			ImageURL:        "./assets/product/fifaXbox.jpg",
			PlatformID:      xboxPlatform.ID,
			Region:          "EU",
			OriginalPrice:   39.99,
			CurrentPrice:    24.99,
			DiscountPercent: 38,
			CashbackAmount:  3.00,
			Likes:           456,
			InStock:         true,
		},
		{
			Name:            "FIFA 23 (EN/PL/CZ/TR) (PC) Origin Key GLOBAL",
			ImageURL:        "./assets/product/fifa.jpg",
			PlatformID:      originPlatform.ID,
			Region:          "GLOBAL",
			OriginalPrice:   39.99,
			CurrentPrice:    24.99,
			DiscountPercent: 38,
			CashbackAmount:  3.00,
			Likes:           456,
			InStock:         true,
		},
		{
			Name:            "FIFA 23 (ENG/PL) (PC) Origin Key EUROPE",
			ImageURL:        "./assets/product/fifa2.jpg",
			PlatformID:      originPlatform.ID,
			Region:          "EU",
			OriginalPrice:   39.99,
			CurrentPrice:    24.99,
			DiscountPercent: 38,
			CashbackAmount:  3.00,
			Likes:           456,
			InStock:         true,
		},
		{
			Name:            "EA SPORTS™ FIFA 23 Ultimate Edition (PC) Steam Key GLOBAL",
			ImageURL:        "./assets/product/fifa.jpg",
			PlatformID:      steamPlatform.ID,
			Region:          "GLOBAL",
			OriginalPrice:   39.99,
			CurrentPrice:    24.99,
			DiscountPercent: 38,
			CashbackAmount:  3.00,
			Likes:           456,
			InStock:         true,
		},
	}

	for _, product := range products {
		if err := DB.Create(&product).Error; err != nil {
			log.Printf("Error creating product %s: %v", product.Name, err)
		}
	}

}
