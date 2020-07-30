/*---
negative:
  phase: resolution
  type: ReferenceError
flags: [module]
---*/
$DONOTEVALUATE();
export {} from './instn-resolve-empty-export_FIXTURE.js';
// instn-resolve-empty-export_FIXTURE.js contains only:
// 0++;