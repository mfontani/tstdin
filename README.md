# tstdin

Timestamp your pipeline.

    $ (echo foo ; sleep 1; echo bar ; sleep 1 ; echo baz) | tstdin
    2031-04-07 14:20:03.952221 00:00:00.000003 00:00:00.000003 foo
    2031-04-07 14:20:04.952120 00:00:00.999902 00:00:00.999898 bar
    2031-04-07 14:20:05.953362 00:00:02.001144 00:00:01.001241 baz

For each line received on STDIN, it outputs:

- (1) the current date/time including microseconds
- (2) hours, minutes, seconds and microseconds since the start
- (3) hours, minutes, seconds and microseconds since the last line was received
- (4) the line that was received

… all separated by spaces.


If the time between the previous line and the current one took too long, the
(3) token will be output with color: red if it took more than a minute, or
yellow if it took more than a second.

Colors are automatically disabled if STDOUT is not a terminal, if `TERM=dumb`
or if either `NO_COLOR` or `TSTDIN_NO_COLOR` are set in the environment, but
see the `-color` and `-no-color` options.

## LICENSE

Copyright 2021 Marco Fontani <MFONTANI@cpan.org>

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
