define(["jquery", "easy-admin"], function ($, ea) {

    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'api/doc.json',
        add_url: 'doc/add.html',
        edit_url: '',
        delete_url: '',
    };

    var Controller = {
        index: function () {
            ea.table.render({
                init: init,
                // toolbar: [
                //     [{
                //         text: '添加',
                //         url: init.add_url,
                //         method: 'open',
                //         auth: 'add',
                //         class: 'layui-btn layui-btn-normal layui-btn-sm',
                //         icon: 'fa fa-plus ',
                //         extend: 'data-full="true"',
                //     }],
                //     'delete'
                // ],
                cols: [[
                    {field: 'category.title', width: 150, title: '分类', sort: true},
                    {field: 'title', minWidth:300, title: '标题', templet:function(d){
                        return '<a href="/doc/id/'+d.id+'">'+ d.title +'</a>'
                      }},
                    {field: 'author', width: 120, title: '作者', search: false},
                    {field: 'create_time', width: 180, title: '创建时间', sort: true, search:false},
                    {
                        width: 150,
                        title: '操作',
                        templet: ea.table.tool,
                        operat: ['edit', 'delete']
                    }
                ]],
            });

            ea.listen();
        },
        add: function () {
            ea.listen();
        },
        edit: function () {
            ea.listen();
        },
    };

    return Controller;
});