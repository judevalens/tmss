build-lib:
	cd lib/mms/cmake-build-debug && cmake --build . && cmake --install .

gen-proto:
	python3 proto/proto.py