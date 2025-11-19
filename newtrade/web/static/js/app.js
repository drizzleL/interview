// 全局变量
let currentUser = 'default_user';
let stocks = [];
let openOrders = [];
let portfolio = null;
let trades = [];

// 页面加载完成后初始化
document.addEventListener('DOMContentLoaded', function () {
    console.log('页面加载完成，开始初始化...');
    loadStocks();
    loadOpenOrders();
    loadPortfolio();
    loadTrades();

    // 设置定时刷新
    setInterval(() => {
        loadStocks();
        loadOpenOrders();
        loadPortfolio();
    }, 5000); // 每5秒刷新一次
});

// 切换用户
function changeUser() {
    const newUser = prompt('请输入新的用户ID:', currentUser);
    if (newUser && newUser.trim()) {
        currentUser = newUser.trim();
        document.getElementById('currentUser').textContent = currentUser;
        loadPortfolio();
        loadTrades();
        showMessage('用户已切换到: ' + currentUser, 'success');
    }
}

// 加载股票数据
async function loadStocks() {
    try {
        console.log('开始加载股票数据...');
        const response = await fetch('/api/stocks');
        console.log('API响应状态:', response.status);

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const result = await response.json();
        console.log('API返回数据:', result);

        if (result.success) {
            stocks = result.data;
            console.log('股票数据加载成功，数量:', stocks.length);
            renderStocks();
            updateOrderSymbols();
        } else {
            console.error('API返回失败:', result);
            showMessage('API返回失败: ' + (result.error || '未知错误'), 'error');
        }
    } catch (error) {
        console.error('加载股票数据失败:', error);
        showMessage('加载股票数据失败: ' + error.message, 'error');
    }
}

// 渲染股票列表
function renderStocks() {
    const stocksGrid = document.getElementById('stocksGrid');
    if (!stocksGrid) {
        console.error('找不到stocksGrid元素');
        return;
    }

    stocksGrid.innerHTML = '';

    // 按symbol排序，保证稳定展示顺序
    const sorted = [...stocks].sort((a, b) => a.symbol.localeCompare(b.symbol));

    sorted.forEach(stock => {
        const stockCard = document.createElement('div');
        stockCard.className = 'stock-card';

        const changeClass = stock.change >= 0 ? 'positive' : 'negative';
        const changeSymbol = stock.change >= 0 ? '+' : '';

        stockCard.innerHTML = `
            <h3>${stock.symbol} - ${stock.name}</h3>
            <div class="stock-price">$${stock.price.toFixed(2)}</div>
            <div class="stock-change ${changeClass}">
                ${changeSymbol}${stock.change.toFixed(2)}%
            </div>
            <div style="margin-top: 10px; font-size: 0.8rem; opacity: 0.8;">
                成交量: ${stock.volume.toLocaleString()}<br>
                市值: $${(stock.marketCap / 1e9).toFixed(1)}B
            </div>
        `;

        // 点击股票卡片自动填充交易表单
        stockCard.addEventListener('click', () => {
            const symbolSelect = document.getElementById('orderSymbol');
            const priceInput = document.getElementById('orderPrice');
            if (symbolSelect && priceInput) {
                symbolSelect.value = stock.symbol;
                priceInput.value = stock.price.toFixed(2);
            }
        });

        stocksGrid.appendChild(stockCard);
    });
}

// 更新订单股票选择器
function updateOrderSymbols() {
    const select = document.getElementById('orderSymbol');
    if (!select) {
        console.error('找不到orderSymbol选择器');
        return;
    }

    const selectedValue = select.value; // 保存当前选中的值

    select.innerHTML = '<option value="">选择股票</option>';

    const sorted = [...stocks].sort((a, b) => a.symbol.localeCompare(b.symbol));

    sorted.forEach(stock => {
        const option = document.createElement('option');
        option.value = stock.symbol;
        option.textContent = `${stock.symbol} - ${stock.name}`;
        select.appendChild(option);
    });

    // 尝试恢复之前的选中状态
    if (selectedValue) {
        select.value = selectedValue;
    }
}

// 下单
async function placeOrder() {
    const symbol = document.getElementById('orderSymbol').value;
    const type = document.getElementById('orderType').value;
    const quantity = parseInt(document.getElementById('orderQuantity').value);
    const price = parseFloat(document.getElementById('orderPrice').value);

    if (!symbol || !type || !quantity || !price) {
        showMessage('请填写完整的订单信息', 'error');
        return;
    }

    if (quantity <= 0 || price <= 0) {
        showMessage('数量和价格必须大于0', 'error');
        return;
    }

    try {
        const response = await fetch('/api/place-order', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                symbol,
                type,
                quantity,
                price,
                userId: currentUser
            })
        });

        const result = await response.json();

        if (result.success) {
            showMessage('订单提交成功！', 'success');
            // 清空表单
            document.getElementById('orderQuantity').value = '';
            document.getElementById('orderPrice').value = '';
            // 刷新数据
            loadOpenOrders();
            loadPortfolio();
            loadTrades();
        } else {
            showMessage('订单提交失败: ' + result.error, 'error');
        }
    } catch (error) {
        console.error('下单失败:', error);
        showMessage('下单失败，请重试', 'error');
    }
}

// 取消订单
async function cancelOrder(orderId) {
    if (!confirm('确定要取消这个订单吗？')) {
        return;
    }

    try {
        const response = await fetch('/api/cancel-order', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ orderId })
        });

        const result = await response.json();

        if (result.success) {
            showMessage('订单已取消', 'success');
            loadOpenOrders();
            loadPortfolio();
        } else {
            showMessage('取消订单失败: ' + result.error, 'error');
        }
    } catch (error) {
        console.error('取消订单失败:', error);
        showMessage('取消订单失败，请重试', 'error');
    }
}

// 加载挂单
async function loadOpenOrders() {
    try {
        const response = await fetch('/api/open-orders');
        const result = await response.json();

        if (result.success) {
            openOrders = result.data;
            renderOpenOrders();
        }
    } catch (error) {
        console.error('加载挂单失败:', error);
    }
}

// 渲染挂单列表
function renderOpenOrders() {
    const table = document.getElementById('openOrdersTable');
    if (!table) {
        console.error('找不到openOrdersTable元素');
        return;
    }

    if (openOrders.length === 0) {
        table.innerHTML = '<p style="text-align: center; color: #666;">暂无挂单</p>';
        return;
    }

    let html = `
        <table>
            <thead>
                <tr>
                    <th>股票</th>
                    <th>类型</th>
                    <th>数量</th>
                    <th>价格</th>
                    <th>用户</th>
                    <th>状态</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody>
    `;

    openOrders.forEach(order => {
        const typeClass = order.Type === 'BUY' ? 'positive' : 'negative';
        const typeText = order.Type === 'BUY' ? '买入' : '卖出';

        html += `
            <tr>
                <td><strong>${order.Symbol}</strong></td>
                <td><span class="${typeClass}">${typeText}</span></td>
                <td>${order.Quantity}</td>
                <td>$${order.Price.toFixed(2)}</td>
                <td>${order.UserID}</td>
                <td>${order.Status}</td>
                <td>
                    ${order.UserID === currentUser ?
                `<button onclick="cancelOrder('${order.ID}')" style="background: #ef4444; color: white; border: none; padding: 5px 10px; border-radius: 5px; cursor: pointer;">取消</button>` :
                '-'
            }
                </td>
            </tr>
        `;
    });

    html += '</tbody></table>';
    table.innerHTML = html;
}

// 加载投资组合
async function loadPortfolio() {
    try {
        const response = await fetch(`/api/portfolio?userId=${currentUser}`);
        const result = await response.json();

        if (result.success) {
            portfolio = result.data;
            renderPortfolio();
        }
    } catch (error) {
        console.error('加载投资组合失败:', error);
    }
}

// 渲染投资组合
function renderPortfolio() {
    const portfolioInfo = document.getElementById('portfolioInfo');
    if (!portfolioInfo) {
        console.error('找不到portfolioInfo元素');
        return;
    }

    if (!portfolio) {
        portfolioInfo.innerHTML = '<p style="text-align: center; color: #666;">暂无投资组合信息</p>';
        return;
    }

    let html = `
        <div class="portfolio-summary">
            <h3>总资产概览</h3>
            <div class="total-value">$${portfolio.TotalValue.toLocaleString('en-US', { minimumFractionDigits: 2 })}</div>
            <div>现金: $${portfolio.Cash.toLocaleString('en-US', { minimumFractionDigits: 2 })}</div>
        </div>
    `;

    if (Object.keys(portfolio.Holdings).length > 0) {
        html += '<div class="holdings-list">';
        html += '<h4>持仓详情</h4>';

        Object.values(portfolio.Holdings).forEach(holding => {
            const stock = stocks.find(s => s.symbol === holding.Symbol);
            const currentPrice = stock ? stock.price : 0;
            const marketValue = holding.Quantity * currentPrice;
            const unrealizedPL = marketValue - holding.TotalCost;
            const plClass = unrealizedPL >= 0 ? 'positive' : 'negative';

            html += `
                <div class="holding-item">
                    <div style="display: flex; justify-content: space-between; align-items: center;">
                        <div>
                            <strong>${holding.Symbol}</strong><br>
                            数量: ${holding.Quantity} | 成本: $${holding.AvgPrice.toFixed(2)}
                        </div>
                        <div style="text-align: right;">
                            <div>市值: $${marketValue.toFixed(2)}</div>
                            <div class="${plClass}">
                                盈亏: $${unrealizedPL.toFixed(2)}
                            </div>
                        </div>
                    </div>
                </div>
            `;
        });

        html += '</div>';
    } else {
        html += '<p style="text-align: center; color: #666;">暂无持仓</p>';
    }

    portfolioInfo.innerHTML = html;
}

// 加载交易历史
async function loadTrades() {
    try {
        const response = await fetch(`/api/trades?userId=${currentUser}`);
        const result = await response.json();

        if (result.success) {
            trades = result.data;
            renderTrades();
        }
    } catch (error) {
        console.error('加载交易历史失败:', error);
    }
}

// 渲染交易历史
function renderTrades() {
    const table = document.getElementById('tradesTable');
    if (!table) {
        console.error('找不到tradesTable元素');
        return;
    }

    if (trades.length === 0) {
        table.innerHTML = '<p style="text-align: center; color: #666;">暂无交易记录</p>';
        return;
    }

    let html = `
        <table>
            <thead>
                <tr>
                    <th>时间</th>
                    <th>股票</th>
                    <th>类型</th>
                    <th>数量</th>
                    <th>价格</th>
                    <th>金额</th>
                </tr>
            </thead>
            <tbody>
    `;

    trades.forEach(trade => {
        const typeClass = trade.Type === 'BUY' ? 'positive' : 'negative';
        const typeText = trade.Type === 'BUY' ? '买入' : '卖出';
        const time = new Date(trade.Timestamp).toLocaleString('zh-CN');

        html += `
            <tr>
                <td>${time}</td>
                <td><strong>${trade.Symbol}</strong></td>
                <td><span class="${typeClass}">${typeText}</span></td>
                <td>${trade.Quantity}</td>
                <td>$${trade.Price.toFixed(2)}</td>
                <td>$${trade.Total.toFixed(2)}</td>
            </tr>
        `;
    });

    html += '</tbody></table>';
    table.innerHTML = html;
}

// 显示消息提示
function showMessage(message, type = 'success') {
    const messageEl = document.getElementById('message');
    if (!messageEl) {
        console.error('找不到message元素');
        return;
    }

    messageEl.textContent = message;
    messageEl.className = `message ${type}`;

    // 3秒后自动隐藏
    setTimeout(() => {
        messageEl.className = 'message hidden';
    }, 3000);
}
