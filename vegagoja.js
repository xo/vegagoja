Date.prototype.setYear = Date.prototype.setFullYear;

function logger(logf) {
  return {
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
}

function loader(logf, loadf) {
  return {
    load(name, res) {
      var s = "";
      try {
        s = loadf(name, res.response);
      } catch (e) {
        logf(["LOAD ERROR", e]);
      }
      return s;
    },
  };
}

function parse(spec) {
  if (typeof spec == "object") {
    return spec;
  }
  return JSON.parse(spec);
}

function vega_version() {
  return vega.version;
}

function lite_version() {
  return vegaLite.version;
}

function render(logf, spec, loadf, cb, errcb) {
  try {
    var runtime = vega.parse(parse(spec));
    var view = new vega.View(runtime, {
      logLevel: vega.Debug,
      logger: logger(logf),
      loader: loader(logf, loadf),
    });
    view.toSVG().then(cb);
  } catch (e) {
    logf(["RENDER ERROR", e]);
    errcb(e);
  } finally {
    if (view) {
      view.finalize();
    }
  }
}

function compile(logf, spec) {
  var res = {};
  try {
    res = vegaLite.compile(parse(spec), {
      logger: logger(logf),
    });
  } catch (e) {
    logf(["COMPILE ERROR", e]);
    throw e;
  }
  return JSON.stringify(res);
}
