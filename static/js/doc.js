define(["jquery", "easy-admin"], function ($, ea) {

    var init = {
        table_elem: '#currentTable',
        table_render_id: 'currentTableRenderId',
        index_url: 'doc/index',
        add_url: 'doc/add',
        edit_url: 'doc/edit',
        delete_url: 'doc/delete',
    };

    var Controller = {
        index: function () {
            ea.table.render({
                init: init,
                cols: [[
                    {field: 'id', width: 80, title: 'ID', sort: true},
                    {field: 'category', width: 150, title: '分类', sort: true},
                    {field: 'title', minWidth:300, title: '标题', templet:function(d){
                        return '<a href="/doc/view/'+d.id+'">'+ d.title +'</a>'
                      }},
                    {field: 'username', width: 120, title: '作者', search: false},
                    {field: 'create_at', width: 180, title: '创建时间', sort: true, search:false},
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