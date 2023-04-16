let s:binPath = expand('<sfile>:p:h')

" s:osDetect {{{
function! s:osDetect() abort
  if has('win64')
      let os = 'windows'
      let bin = 'vasr_windows.exe'
  else
      if has('mac')
        let os = 'darwin'
        let bin = 'vasr_darwin'
      else
        let os = 'linux'
        let bin = 'vasr_linux'
      endif
  endif

  return [l:os, l:bin]
endfunction
" }}}

" NOTE: vim-airline appears to expect the lower-cased
" airline#extensions#foo#bar naming convention.

" airline#extensions#system#init {{{
function! airline#extensions#system#init(ext) abort
  call airline#parts#define_raw('system', '%{airline#extensions#system#get()}')
  call a:ext.add_statusline_func('airline#extensions#system#apply')
endfunction
" }}}

" airline#extensions#system#apply {{{
function! airline#extensions#system#apply(...) abort
  let w:airline_section_z = '%{airline#extensions#system#get()}'
endfunction
" }}}

" airline#extensions#system#get {{{
function! airline#extensions#system#get() abort
  let host = s:osDetect()
  let os = host[0]
  let bin = host[1]
  let sysinfo = system(s:binPath . '/bin/' . bin)
  return os . ' ' . substitute(sysinfo, '\n\+$', '', '')
endfunction
" }}}

" airline#extensions#systemWtimerfn {{{
function! airline#extensions#system#timerfn(timer) abort
  call airline#update_statusline()
endfunction
" }}}

" g:airline#extensions#system#timer {{{
let g:airline#extensions#system#timer = timer_start(
  \   10000,
  \   'airline#extensions#system#timerfn',
  \   {'repeat':-1}
  \ )
" }}}
