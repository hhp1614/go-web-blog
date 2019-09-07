// 注册表单验证
$('#register-form').validate({
    rules: {
        username: {
            required: true,
            rangelength: [5, 10]
        },
        password: {
            required: true,
            rangelength: [5, 10]
        },
        rePassword: {
            required: true,
            rangelength: [5, 10],
            equalTo: '#register-password'
        }
    },
    messages: {
        username: {
            required: '请输入用户名',
            rangelength: '用户名必须是5-10位'
        },
        password: {
            required: '请输入密码',
            rangelength: '密码必须是5-10位'
        },
        rePassword: {
            required: '请确认密码',
            rangelength: '密码必须是5-10位',
            equalTo: '两次输入的密码必须相等'
        }
    },
    submitHandler(form) {
        const urlStr = '/register';
        console.log('urlStr:', urlStr);
        $(form).ajaxSubmit({
            url: urlStr,
            type: 'post',
            dataType: 'json',
            success(data, status) {
                layer.msg(data.message);
                if (data.code === 1) {
                    setTimeout(() => {
                        window.location.href = '/login';
                    }, 100);
                }
            },
            error(data, status) {
                layer.msg(data.message);
            }
        });
    }
});

// 登录表单验证
$('#login-form').validate({
    rules: {
        username: {
            required: true,
            rangelength: [5, 10]
        },
        password: {
            required: true,
            rangelength: [5, 10]
        }
    },
    messages: {
        username: {
            required: '请输入用户名',
            rangelength: '用户名必须是5-10位'
        },
        password: {
            required: '请输入密码',
            rangelength: '密码必须是5-10位'
        }
    },
    submitHandler(form) {
        const urlStr = '/login';
        console.log('urlStr:', urlStr);
        $(form).ajaxSubmit({
            url: urlStr,
            type: 'post',
            dataType: 'json',
            success(data, status) {
                layer.msg(data.message);
                if (data.code === 1) {
                    setTimeout(() => {
                        window.location.href = '/';
                    }, 100);
                }
            },
            error(data, status) {
                layer.msg(data.message);
            }
        });
    }
});

// 添加文章的表单
$('#write-art-form').validate({
    rules: {
        title: 'required',
        tags: 'required',
        short: {
            required: true,
            minlength: 2
        },
        content: {
            required: true,
            minlength: 2
        }
    },
    messages: {
        title: '请输入标题',
        tags: '请输入标签',
        short: {
            required: '请输入简介',
            minlength: '简介内容最少两个字符'
        },
        content: {
            required: '请输入文章内容',
            minlength: '文章内容最少两个字符'
        }
    },
    submitHandler(form) {
        let urlStr = '/article/add';
        const artId = $('#write-article-id').val();
        if (artId > 0) {
            urlStr = '/article/update'
        }
        console.log('urlStr:', urlStr);
        $(form).ajaxSubmit({
            url: urlStr,
            type: 'post',
            dataType: 'json',
            success(data, status) {
                layer.msg(data.message);
                setTimeout(() => {
                    window.location.href = '/';
                }, 100);
            },
            error(data, status) {
                alert('err:' + data.message + ':' + status);
                layer.msg(data.message);
            }
        });
    }
});
