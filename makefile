mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
xplane_dir := "/Volumes/storage/X-Plane 12"

all: mac win

clean:
	rm -rf dist || true || rm ${xplane_dir}/Resources/plugins/xa-honeycomb/mac.xpl

mac:
	mkdir -p build/xa-honeycomb
	cargo build --release --target aarch64-apple-darwin
	mv target/aarch64-apple-darwin/release/libxa_honeycomb.dylib build/xa-honeycomb/mac_arm.xpl
	cargo build --release --target x86_64-apple-darwin
	mv target/x86_64-apple-darwin/release/libxa_honeycomb.dylib build/xa-honeycomb/mac_amd.xpl
	lipo build/xa-honeycomb/mac_arm.xpl build/xa-honeycomb/mac_amd.xpl -create -output build/xa-honeycomb/mac.xpl
	rm -rf build/xa-honeycomb/mac_arm.xpl build/xa-honeycomb/mac_amd.xpl

dev:
	cargo build
	mv target/debug/libxa_honeycomb.dylib build/xa-honeycomb/mac.xpl
	cp build/xa-honeycomb/mac.xpl ${xplane_dir}/Resources/plugins/xa-honeycomb/mac.xpl

win:
	mkdir -p build/xa-honeycomb
	cargo build --release --target x86_64-pc-windows-gnu
	mv target/x86_64-pc-windows-gnu/release/xa_honeycomb.dll build/xa-honeycomb/win.xpl

# lin:
# 	mkdir -p build/xa-honeycomb
# 	cargo build --release --target x86_64-unknown-linux-gnu
# 	mv target/aarch64-apple-darwin/release/libxa_honeycomb.dylib build/xa-honeycomb/lin.xpl