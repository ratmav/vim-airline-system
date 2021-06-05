let s:binPath = expand('<sfile>:p:h')

" s:osDetect {{{
function! s:osDetect() abort
  if has('win64')
      let os = 'windows'
  else
      if has('mac')
        let os = 'darwin'
      else
        let os = 'linux'
      endif
  endif

  return os
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
  let supported_platforms = ['darwin', 'linux']
  let os = s:osDetect()

  if index(supported_platforms, os) >= 0
    let sysinfo = system(s:binPath . '/bin/' . os)
    return os . ' ' . substitute(sysinfo, '\n\+$', '', '')
  else
    return 'unsupported platform'
  endif
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
