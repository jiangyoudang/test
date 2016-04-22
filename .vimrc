set nocompatible
filetype off

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'VundleVim/Vundle.vim'

Plugin 'fatih/vim-go'
"Plugin 'Shougo/neocomplete'
Plugin 'Valloric/YouCompleteMe'
Plugin 'Valloric/MatchTagAlways'
Plugin 'airblade/vim-gitgutter'
Plugin 'altercation/vim-colors-solarized'


call vundle#end()

"let g:neocomplete#enable_at_startup = 1
"let g:neocomplete#enable_smart_case = 1
if exists('$TMUX')
  let &t_SI = "\<Esc>Ptmux;\<Esc>\<Esc>]50;CursorShape=1\x7\<Esc>\\"
  let &t_EI = "\<Esc>Ptmux;\<Esc>\<Esc>]50;CursorShape=0\x7\<Esc>\\"
else
  let &t_SI = "\<Esc>]50;CursorShape=1\x7"
  let &t_EI = "\<Esc>]50;CursorShape=0\x7"
endif

set nu
syntax on
set background=light
colorscheme solarized
set laststatus=2
set showcmd
set ruler
" tabs = 2
set ts=2
" expand tab into space
set expandtab
" auto indent
set autoindent
set cursorline

" >> or << by 2 spaces
set shiftwidth=2

filetype plugin indent on
