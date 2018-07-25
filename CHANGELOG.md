## 0.5.0 (Unreleased)

  * Created metric multiemitter to allow multiple sources.
  * Removed aws signing client (use aws-signing instead).
  * Rebuilt metrics to be more generic.

## 0.4.2 (Jul 04, 2018)

  * Expanded metrics to include dimensions.
  * Added SetResolution to Point.
  * Defaulted Point to low resolution.

## 0.4.1 (Jul 03, 2018)

  * Added DurationPoint for cloudwatch metrics.

## 0.4.0 (Jul 02, 2018)

  * Upgraded aws-sdk to v2.

## 0.3.2 (Jun 22, 2018)

  * Added relay `NewMultiEmitter`.
  * Added relay `NewElasticsearchEmitter`.

## 0.3.1 (Jun 15, 2018)

  * Added configuration option to results Saver `AlwaysSave`.

## 0.3.0 (Jun 04, 2018)

  * Converted relay emitter to on-the-fly.

## 0.2.3 (Jun 01, 2018)

  * Added noop interfaces

## 0.2.2 (Jun 01, 2018)

  * Added extensibility of contextual interfaces. 

## 0.2.1 (May 29, 2018)

  * Configured `CountPoint` with correct `Unit` of measurement.

## 0.2.0 (May 29, 2018)

  * Rebuilt platform with better interface support using `context`.

## 0.1.6 (May 25, 2018)

  * Added cloudwatch metrics emitter.

## 0.1.5 (May 21, 2018)

  * Fixed bug where trying to emit 0 items errors.

## 0.1.4 (May 21, 2018)

  * Added cloudwatch PutMetrics.

## 0.1.3 (May 14, 2018)

  * Added dynamodb Get.
  * Added dynamodb Put.

## 0.1.2 (May 02, 2018)

  * Emitting kinesis records in batches to prevent exceeding AWS API limits.

## 0.1.1 (May 01, 2018)

  * Added `parallel` package with `All`.

## 0.1.0 (May 01, 2018)

  * Added kinesis input event.
  * Added kinesis Emit().
