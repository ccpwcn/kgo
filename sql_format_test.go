package kgo

import (
	"fmt"
	"testing"
)

func TestFormatSQLWithDefaultOptions(t *testing.T) {
	type args struct {
		sql string
	}
	type testCase struct {
		name string
		args args
		want bool
	}
	tests := []testCase{
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "SELECT * FROM my_table_name",
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "SELECT COUNT(*) from sys_user",
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "select * from users where id = 1",
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "UPDATE users SET last_login = NOW() WHERE  id = 123;",
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "SELECT * FROM users WHERE age > 18 AND name LIKE '%John%' ORDER BY age DESC;",
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "SELECT * FROM a WHERE x IN (SELECT y FROM b WHERE z IN (SELECT w FROM c WHERE v = 1))",
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "SELECT u.id,u.name,(SELECT COUNT(*) FROM orders o WHERE o.user_id=u.id) AS order_count FROM users u WHERE u.active=true AND (u.created_at>'2023-01-01' OR u.email LIKE '%@example.com') ORDER BY u.name",
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql: "SELECT u.id, u.name, COUNT(o.id) AS order_count FROM users u LEFT JOIN orders o ON u.id = o.user_id WHERE u.active = TRUE AND u.created_at > '2023-01-01' GROUP BY u.id, u.name HAVING COUNT(o.id) > 0 ORDER BY u.name;",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatSQL(tt.args.sql)
			fmt.Println(got)
			if !tt.want {
				t.Errorf("FormatSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatSQLWithSomeOptions(t *testing.T) {
	type args struct {
		sql     string
		options []SQLFormatOption
	}
	type testCase struct {
		name string
		args args
		want bool
	}
	opts := []SQLFormatOption{
		WithUpperCaseKeywords(true), // 保留关键字大小写
		WithIndent("  "),            // 使用2空格缩进
		WithLineWidth(100),          // 换行
	}
	tests := []testCase{
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql:     "SELECT * FROM my_table_name",
				options: opts,
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql:     "select * from users where id = 1",
				options: opts,
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql:     "UPDATE users SET last_login = NOW() WHERE  id = 123;",
				options: opts,
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql:     "SELECT * FROM users WHERE age > 18 AND name LIKE '%John%' ORDER BY age DESC;",
				options: opts,
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql:     "SELECT * FROM a WHERE x IN (SELECT y FROM b WHERE z IN (SELECT w FROM c WHERE v = 1))",
				options: opts,
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql:     "SELECT u.id,u.name,(SELECT COUNT(*) FROM orders o WHERE o.user_id=u.id) AS order_count FROM users u WHERE u.active=true AND (u.created_at>'2023-01-01' OR u.email LIKE '%@example.com') ORDER BY u.name",
				options: opts,
			},
			want: true,
		},
		{
			name: "formatSQLWithDefaultOptions",
			args: args{
				sql:     "SELECT u.id, u.name, COUNT(o.id) AS order_count FROM users u LEFT JOIN orders o ON u.id = o.user_id WHERE u.active = TRUE AND u.created_at > '2023-01-01' GROUP BY u.id, u.name HAVING COUNT(o.id) > 0 ORDER BY u.name;",
				options: opts,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatSQL(tt.args.sql, tt.args.options...)
			fmt.Println(got)
			if !tt.want {
				t.Errorf("FormatSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
