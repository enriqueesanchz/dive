package export

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/wagoodman/dive/dive/image/docker"
)

func Test_Export(t *testing.T) {
	result := docker.TestAnalysisFromArchive(t, "../../.data/test-docker-image.tar")

	export := NewExport(result)
	payload, err := export.Marshal()
	if err != nil {
		t.Errorf("Test_Export: unable to export analysis: %v", err)
	}

	expectedResult := `{
  "layer": [
    {
      "index": 0,
      "id": "blobs",
      "digestId": "sha256:155830e8292e8dbeb0d3b0612276ef41cb48680a45e09dda90df765eb7538eb8",
      "sizeBytes": 1483227,
      "command": "",
      "fileList": [
        {
          "Path": "1871059774abe6914075e4a919b778fa1561f577d620ae52438a9635e6241936/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "1871059774abe6914075e4a919b778fa1561f577d620ae52438a9635e6241936/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "1871059774abe6914075e4a919b778fa1561f577d620ae52438a9635e6241936/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 8192,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "1871059774abe6914075e4a919b778fa1561f577d620ae52438a9635e6241936",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "28cfe03618aa2e914e81fdd90345245c15f4478e35252c06ca52d238fd3cc694/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "28cfe03618aa2e914e81fdd90345245c15f4478e35252c06ca52d238fd3cc694/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 406,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "28cfe03618aa2e914e81fdd90345245c15f4478e35252c06ca52d238fd3cc694/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 1369600,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "28cfe03618aa2e914e81fdd90345245c15f4478e35252c06ca52d238fd3cc694",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "3d4ad907517a021d86a4102d2764ad2161e4818bbd144e41d019bfc955434181/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "3d4ad907517a021d86a4102d2764ad2161e4818bbd144e41d019bfc955434181/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "3d4ad907517a021d86a4102d2764ad2161e4818bbd144e41d019bfc955434181/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 8704,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "3d4ad907517a021d86a4102d2764ad2161e4818bbd144e41d019bfc955434181",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "461885fc22589158dee3c5b9f01cc41c87805439f58b4399d733b51aa305cbf9/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "461885fc22589158dee3c5b9f01cc41c87805439f58b4399d733b51aa305cbf9/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "461885fc22589158dee3c5b9f01cc41c87805439f58b4399d733b51aa305cbf9/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 9728,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "461885fc22589158dee3c5b9f01cc41c87805439f58b4399d733b51aa305cbf9",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "49fe2a475548bfa4d493fc796fce41f30704e3d4cbff3e45dd3e06f463236d1d/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "49fe2a475548bfa4d493fc796fce41f30704e3d4cbff3e45dd3e06f463236d1d/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "49fe2a475548bfa4d493fc796fce41f30704e3d4cbff3e45dd3e06f463236d1d/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3072,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "49fe2a475548bfa4d493fc796fce41f30704e3d4cbff3e45dd3e06f463236d1d",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "5eca617bdc3bc06134fe957a30da4c57adb7c340a6d749c8edc4c15861c928d7/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "5eca617bdc3bc06134fe957a30da4c57adb7c340a6d749c8edc4c15861c928d7/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "5eca617bdc3bc06134fe957a30da4c57adb7c340a6d749c8edc4c15861c928d7/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 9216,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "5eca617bdc3bc06134fe957a30da4c57adb7c340a6d749c8edc4c15861c928d7",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "75ae28b8ebf89b203319069ddcbecdf5c3838503bbebc0d0d85e4b8589c5de3a.json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3994,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "80cd2ca1ffc89962b9349c80280c2bc551acbd11e09b16badb0669f8e2369020/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "80cd2ca1ffc89962b9349c80280c2bc551acbd11e09b16badb0669f8e2369020/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "80cd2ca1ffc89962b9349c80280c2bc551acbd11e09b16badb0669f8e2369020/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 9216,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "80cd2ca1ffc89962b9349c80280c2bc551acbd11e09b16badb0669f8e2369020",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "81b1b002d4b4c1325a9cad9990b5277e7f29f79e0f24582344c0891178f95905/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "81b1b002d4b4c1325a9cad9990b5277e7f29f79e0f24582344c0891178f95905/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "81b1b002d4b4c1325a9cad9990b5277e7f29f79e0f24582344c0891178f95905/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 9216,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "81b1b002d4b4c1325a9cad9990b5277e7f29f79e0f24582344c0891178f95905",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "a10327f68ffed4afcba78919052809a8f774978a6b87fc117d39c53c4842f72c/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "a10327f68ffed4afcba78919052809a8f774978a6b87fc117d39c53c4842f72c/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "a10327f68ffed4afcba78919052809a8f774978a6b87fc117d39c53c4842f72c/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 8704,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "a10327f68ffed4afcba78919052809a8f774978a6b87fc117d39c53c4842f72c",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "aad36d0b05e71c7e6d4dfe0ca9ed6be89e2e0d8995dafe83438299a314e91071/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "aad36d0b05e71c7e6d4dfe0ca9ed6be89e2e0d8995dafe83438299a314e91071/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "aad36d0b05e71c7e6d4dfe0ca9ed6be89e2e0d8995dafe83438299a314e91071/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 5632,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "aad36d0b05e71c7e6d4dfe0ca9ed6be89e2e0d8995dafe83438299a314e91071",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "c99e2f8d3f6282668f0d30dc1db5e67a51d7a1dcd7ff6ddfa0f90760836778ec/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "c99e2f8d3f6282668f0d30dc1db5e67a51d7a1dcd7ff6ddfa0f90760836778ec/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "c99e2f8d3f6282668f0d30dc1db5e67a51d7a1dcd7ff6ddfa0f90760836778ec/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 9216,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "c99e2f8d3f6282668f0d30dc1db5e67a51d7a1dcd7ff6ddfa0f90760836778ec",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "cfb35bb5c127d848739be5ca726057e6e2c77b2849f588e7aebb642c0d3d4b7b/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "cfb35bb5c127d848739be5ca726057e6e2c77b2849f588e7aebb642c0d3d4b7b/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 1239,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "cfb35bb5c127d848739be5ca726057e6e2c77b2849f588e7aebb642c0d3d4b7b/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 8704,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "cfb35bb5c127d848739be5ca726057e6e2c77b2849f588e7aebb642c0d3d4b7b",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "f07c3eb887572395408f8e11a07af945e4da5f02b3188bb06b93fad713ca0b99/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "f07c3eb887572395408f8e11a07af945e4da5f02b3188bb06b93fad713ca0b99/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "f07c3eb887572395408f8e11a07af945e4da5f02b3188bb06b93fad713ca0b99/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 9216,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "f07c3eb887572395408f8e11a07af945e4da5f02b3188bb06b93fad713ca0b99",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "f2fc54e25cb7966dc9732ec671a77a1c5c104e732bd15ad44a2dc1ac42368f84/VERSION",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 3,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "f2fc54e25cb7966dc9732ec671a77a1c5c104e732bd15ad44a2dc1ac42368f84/json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 482,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "f2fc54e25cb7966dc9732ec671a77a1c5c104e732bd15ad44a2dc1ac42368f84/layer.tar",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 2048,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "f2fc54e25cb7966dc9732ec671a77a1c5c104e732bd15ad44a2dc1ac42368f84",
          "TypeFlag": 53,
          "Linkname": "",
          "Size": 0,
          "Mode": 2147484141,
          "Uid": 0,
          "Gid": 0,
          "IsDir": true
        },
        {
          "Path": "manifest.json",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 1206,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        },
        {
          "Path": "repositories",
          "TypeFlag": 48,
          "Linkname": "",
          "Size": 92,
          "Mode": 420,
          "Uid": 0,
          "Gid": 0,
          "IsDir": false
        }
      ]
    }
  ],
  "image": {
    "sizeBytes": 1483227,
    "inefficientBytes": 0,
    "efficiencyScore": 1,
    "fileReference": []
  }
}`
	actualResult := string(payload)
	if expectedResult != actualResult {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedResult, actualResult, false)

		t.Errorf("Test_Export: unexpected export result:\n%v", dmp.DiffPrettyText(diffs))
	}
}
