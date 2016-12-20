var DiffMatchPatch = require('diff-match-patch');
var dmp = new DiffMatchPatch();

const text1 = "Lorem ipsum dolor.";
const text2 = "Lorem dolor sit amet.";

function createPatch(lastDoc, currDoc) {
    var t0 = process.hrtime()[1];

    var diff = dmp.diff_main(lastDoc, currDoc, false);
    dmp.diff_cleanupSemantic(diff);

    var patch = dmp.patch_make(lastDoc, diff);

    var result = dmp.patch_toText(patch);

    var t1 = process.hrtime()[1];
    console.log((t1 - t0) + 'ns');

    return result;
}

console.log(createPatch(text1, text2));