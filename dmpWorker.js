var fs = require('fs');
var DiffMatchPatch = require('diff-match-patch');
var dmp = new DiffMatchPatch();

const text1 = "Lorem ipsum dolor.";
const text2 = "Lorem dolor sit amet.";

function createPatch(lastDoc, currDoc) {
    var m0 = process.memoryUsage().heapUsed;
    var t0 = process.hrtime()[1];

    var diff = dmp.diff_main(lastDoc, currDoc, false);
    dmp.diff_cleanupSemantic(diff);

    var patch = dmp.patch_make(lastDoc, diff);

    var result = dmp.patch_toText(patch);

    var t1 = process.hrtime()[1];
    console.log((t1 - t0) + ' ns');
    var m1 = process.memoryUsage().heapUsed;
    console.log((m1 - m0) / 1024 + ' KB');

    return result;
}

createPatch(text1, text2)
//console.log(createPatch(text1, text2));
var d1 = fs.readFileSync("speedtest1.txt", 'utf8');
var d2 = fs.readFileSync("speedtest2.txt", 'utf8');
createPatch(d1, d2)
//console.log(createPatch(d1, d2));
