<?php


namespace JSMinGateaway;


class Minifier
{
    protected $ffi;

    function __construct()
    {
        $this->ffi = \FFI::cdef(
            "char* MinifyJS(char* p0, char* p1);",
            __DIR__ . "/bin/gateaway.so"
        );
    }

    public function minify($content)
    {
        $minified_js = $this->ffi->MinifyJS($content, __DIR__);

        return \FFI::string($minified_js);
    }

    public function getDirectory()
    {
        $minified_js = $this->ffi->GetDirectory(__DIR__);
        return \FFI::string($minified_js);
    }
}
