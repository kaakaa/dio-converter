**Move to https://github.com/kaakaa/dio-exporter**

# dio-converter

Convert dio to png

## Requirements
* Docker, Docker Compose
* Go
* Japanese fonts files (if you use Japanese in diagrams)

## Set up

* 1. `git clone https://github.com/kaakaa/dio-converter`
* 2. (If you use Japanese fonts) download Japanese font files
  * Download Japanese fonts file from https://fonts.google.com/?subset=japanese
  * Unzip and move fonts files (*.otf) to `docker-drawio/image-export/standalone/fonts`
* 3. Run `docker-compose up -d`
* 4. Run `go run main.go`
  * find `.dio` files in `./input/` recursively
  * output `.png` files to `./output/`


## TODO
* CLI
  * inputs
  * outputs
  * input types (`.dio`, `.drawio`)
  * output types (png, svg, pdf)
