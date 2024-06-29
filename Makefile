default: build

build:
	@fyne build

package:
	@fyne package --name="Kitchen Clock" --icon="icon.png"
