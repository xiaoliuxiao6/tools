#!/bin/bash

# 天蓝
echoCyan(){
    echo -e "\033[36m$*\033[0m"
}
################################################################################################
echoCyan "设置 ls 输出时间格式"
################################################################################################
i="export TIME_STYLE='+%Y-%m-%d %H:%M:%S'"
echo ${i} >> ~/.bashrc

##############################################################################
echoCyan "VIM 设置"
##############################################################################
cat <<\EOF > ~/.vimrc
" 启用代码折叠
set foldmethod=indent
set foldlevel=99
nnoremap <space> za

" 支持 UTF-8 编码
set encoding=utf-8

" 按 F5 自动执行
map <F5> :call CompileRunGcc()<CR>
func! CompileRunGcc()
        exec "w"
        if &filetype == 'c'
                exec "!g++ % -o %<"
                exec "!time ./%<"
        elseif &filetype == 'cpp'
                exec "!g++ % -o %<"
                exec "!time ./%<"
        elseif &filetype == 'java'
                exec "!javac %"
                exec "!time java %<"
        elseif &filetype == 'sh'
                :!time bash %
        elseif &filetype == 'python'
                exec "!clear"
                exec "!time python3 %"
        elseif &filetype == 'html'
                exec "!firefox % &"
        elseif &filetype == 'go'
                " exec "!go build %<"
                exec "!time go run %"
        elseif &filetype == 'mkd'
                exec "!~/.vim/markdown.pl % > %.html &"
                exec "!firefox %.html &"
        endif
endfunc

set hlsearch            " 搜索时高亮被找到的文字
set expandtab           " TAB 变空格
set tabstop=4           " 设置 TAB 宽度为4
set ignorecase smartcase    " 搜索时忽略大小写
EOF

##############################################################################
echo "修改 history 行为"
##############################################################################
## 删除原有设置
sed -i '/^HISTSIZE/d' /root/.bashrc
sed -i '/^HISTFILESIZE/d' /root/.bashrc
sed -i '/^HISTTIMEFORMAT/d' /root/.bashrc
sed -i '/^PROMPT_COMMAND/d' /root/.bashrc
sed -i '/^shopt -s/d' /root/.bashrc

## 追加新设置
cat <<\EOF >> /root/.bashrc
## 追加而不是覆盖
shopt -s histappend
## 定义命令输出的行数
HISTSIZE=1000
## 定义最多保留的条数
HISTFILESIZE=2000
## 记录执行命令的时间和用户名
HISTTIMEFORMAT="%Y-%m-%d %H:%M:%S:`whoami` "
## 实时追加，不必等用户退出
PROMPT_COMMAND="history -a"
## 当终端窗口大小改变时，确保显示得到更新
shopt -s checkwinsize
EOF

##############################################################################
echo "修改 Ulimit（重启生效）"
##############################################################################
sed -i '/DefaultLimitNOFILE/d' /etc/systemd/system.conf
sed -i '/DefaultLimitNOFILE/d' /etc/systemd/user.conf
echo 'DefaultLimitNOFILE=1048576' >> /etc/systemd/system.conf
echo 'DefaultLimitNOFILE=1048576' >> /etc/systemd/user.conf

sed -i "/nofile/d" /etc/security/limits.conf
echo "root soft nofile 1048576" >> /etc/security/limits.conf
echo "root hard nofile 1048576" >> /etc/security/limits.conf

sed -i "/nproc/d" /etc/security/limits.conf
echo 'root soft nproc 1030567' >> /etc/security/limits.conf
echo 'root hard nproc 1030567' >> /etc/security/limits.conf

ulimit -u 1048576
ulimit -n 1048576
sysctl -p

##############################################################################
echo "取消开机30秒等待"
##############################################################################
sed -i '/GRUB_TIMEOUT=/d' /etc/default/grub
sed -i '/GRUB_RECORDFAIL_TIMEOUT=/d' /etc/default/grub
echo 'GRUB_TIMEOUT=3' >> /etc/default/grub
echo 'GRUB_RECORDFAIL_TIMEOUT=3' >> /etc/default/grub

update-grub

# ##############################################################################
# echo "允许 root 用户使用 SSH 登录"
# ##############################################################################
# echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config
# systemctl restart sshd
# systemctl status sshd


##############################################################################
echo "禁用每日更新"
##############################################################################
sudo systemctl stop apt-daily.service
sudo systemctl stop apt-daily.timer
sudo systemctl stop apt-daily-upgrade.service
sudo systemctl stop apt-daily-upgrade.timer
sudo systemctl disable apt-daily.service
sudo systemctl disable apt-daily.timer
sudo systemctl disable apt-daily-upgrade.service
sudo systemctl disable apt-daily-upgrade.timer

