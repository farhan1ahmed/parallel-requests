# Parallel Requests Tool
The parallel requests tools is capable of sending http calls in parallel. The tool takes in a flag argument of
parallel which determines the number of concurrent go-routines to start. The flag has a default value of 10 which prevents
exhausting machine's resources.

In order to use the tool, build the myhttp.go file
``go build myhttp.go``
This will build the executable file of the code.

The executable can be used to make parallel requests and print md5 Hashes of the response bodies. Here are few examples:

1. **Command:** ./myhttp http://adjust.com
<br>**Result:**<br>
``http://adjust.com 883516367ac9b28ade46fefd23c07935
   ``
2. **Command:** ./myhttp http://adjust.com http://google.com
   <br>**Result:**<br>
   ``http://adjust.com 883516367ac9b28ade46fefd23c07935``<br>
   ``http://google.com 2a73fad44f7b99cb9690790289fc7164``
3. **Command:** ./myhttp -parallel=4 http://adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://adjust.com https://adjust.com https://google.com
    <br>**Result:**
   ``http://adjust.com 883516367ac9b28ade46fefd23c07935``<br>
   ``http://adjust.com 883516367ac9b28ade46fefd23c07935``<br>
   ``https://adjust.com 883516367ac9b28ade46fefd23c07935``<br>
   ``https://google.com 59074c20c3c6db1a844b8212d57a3082``<br>
   ``http://google.com faf78e080f20e4a697e22ce664adead8``<br>
   ``http://yandex.com b64620794950f4d56dffeabfdce32e93``<br>
   ``http://yahoo.com 1e2c43d9cbbacde62a64403c954c7e85``<br>
   ``http://twitter.com 46a07273526c61d0d1c6629357957dce``<br>
   ``http://facebook.com 1766eee06b748fbd3eabd43471086360``<br>