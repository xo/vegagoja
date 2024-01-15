function version() {
  return vega.version;
}

function render(logf, loadf, cb, spec, data) {
  const logger = {
    level(_) {},
    error() {
      logf(["ERROR", ...arguments]);
    },
    warn() {
      logf(["WARN", ...arguments]);
    },
    info() {
      logf(["INFO", ...arguments]);
    },
    debug() {
      logf(["DEBUG", ...arguments]);
    },
  };
  const loader = {
    load(name) {
      var s = "";
      try {
        s = loadf(name);
      } catch (e) {
        logf(["LOAD ERROR", e]);
      }
      //logf(["result", s]);
      return "";
      //return JSON.parse(s);
    },
  };
  try {
    var s = JSON.parse(spec);
    var runtime = vega.parse(s);
    var view = new vega.View(runtime, {
      loader: loader,
    });
    view.logger(logger).toSVG().then(cb);
  } catch (e) {
    throw e;
  } finally {
    if (view) {
      view.finalize();
    }
  }
}
